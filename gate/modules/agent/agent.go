package agent

import (
	"context"
	"eastv2/pb"
	"sync"

	"github.com/tnnmigga/corev2/conc"
	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/message"
	"github.com/tnnmigga/corev2/system"
	"github.com/tnnmigga/corev2/utils"

	"sync/atomic"
	"time"
)

const (
	MaxSendMQLen = 64
	MaxAliveTime = 10 * time.Second
)

const (
	AgentStateNone = iota
	AgentStateRun
	AgentStateWait
	AgentStateDead
)

type AgentManager struct {
	iface.IModule
	agents map[uint64]*Agent
	rw     sync.RWMutex
}

func NewAgentManager(m iface.IModule) *AgentManager {
	am := &AgentManager{
		IModule: m,
		agents:  map[uint64]*Agent{},
	}
	conc.Go(am.clean)
	return am
}

func (am *AgentManager) AddAgent(agent *Agent) {
	if agent.userID == 0 {
		log.Errorf("agent uid is 0")
		return
	}
	am.rw.Lock()
	am.agents[agent.userID] = agent
	am.rw.Unlock()
}

func (am *AgentManager) RemoveAgent(uid uint64) {
	am.rw.Lock()
	delete(am.agents, uid)
	am.rw.Unlock()
}

func (am *AgentManager) GetAgent(uid uint64) *Agent {
	am.rw.RLock()
	agent := am.agents[uid]
	am.rw.RUnlock()
	return agent
}

func (am *AgentManager) OnConnect(conn Conn) {
	// resp, err := message.AnyRPC[pb.TokenAuthResp](define.ServDoor, &pb.TokenAuthReq{})
	// if err != nil {
	// 	log.Errorf("TokenAuthReq return error %v", err)
	// 	conn.Close()
	// 	return
	// }
	agent := NewAgent()
	agent.conn = conn
	agent.servID = 1001
	agent.userID = 1
	am.AddAgent(agent)
	agent.Run()
}

func (am *AgentManager) OnError(err error) {
	log.Errorf("agent manager error %v", err)
}

func (am *AgentManager) clean() {
	ticker := time.NewTicker(MaxAliveTime)
	for {
		select {
		case <-ticker.C:
			am.cleanDeadAgent()
		case <-system.RootCtx().Done():
			return
		}
	}
}

func (am *AgentManager) cleanDeadAgent() {
	am.rw.RLock()
	defer am.rw.RUnlock()
	nowNs := utils.NowNs()
	for uid, agent := range am.agents {
		state := atomic.LoadInt32(&agent.state)
		if state == AgentStateDead {
			delete(am.agents, uid)
			continue
		}
		if state != AgentStateWait {
			continue
		}
		if agent.waitNs+MaxAliveTime > nowNs {
			continue
		}
		if !atomic.CompareAndSwapInt32(&agent.state, AgentStateWait, AgentStateDead) {
			continue
		}
		delete(am.agents, uid)
	}
}

type IAgent interface {
	OnMessage([]byte)
	OnReconnect()
	OnReadError(error)
}

type Agent struct {
	userID  uint64
	servID  uint32
	conn    Conn
	sendMQ  chan []byte
	state   int32
	failMsg []byte // 记录发送失败的消息等待重发
	waitNs  time.Duration
	cancel  func()
}

func NewAgent() *Agent {
	agent := &Agent{
		sendMQ: make(chan []byte, MaxSendMQLen),
	}
	return agent
}

func (a *Agent) OnMessage(data []byte) {
	message.Cast(a.servID, &pb.C2SMsg{
		UserID: a.userID,
		Body:   data,
	})
}

func (a *Agent) OnError(err error) {
	if atomic.CompareAndSwapInt32(&a.state, AgentStateRun, AgentStateWait) {
		a.waitNs = utils.NowNs()
		log.Debugf("agent read error %v", err)
		a.conn.Close()
		a.cancel()
	}
}

func (a *Agent) OnReconnect() {
	if atomic.CompareAndSwapInt32(&a.state, AgentStateWait, AgentStateRun) {
		log.Debugf("agent reconnect")
	}
	if a.failMsg != nil {
		if err := a.conn.Write(a.failMsg); err != nil {
			a.OnError(err)
			return
		}
	}
	a.Run()
}

func (a *Agent) Run() {
	a.state = AgentStateRun
	// 启动agent读写goroutine
	ctx, cancel := context.WithCancel(context.Background())
	if a.cancel != nil {
		log.Warnf("cancel is exist")
	}
	a.cancel = cancel
	a.readLoop(ctx)
	a.writeLoop(ctx)
}

func (a *Agent) readLoop(ctx context.Context) {
	conc.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				data, err := a.conn.Read(MaxAliveTime)
				if err != nil {
					a.OnError(err)
					return
				}
				if len(data) == 0 {
					continue // 心跳
				}
				a.OnMessage(data)
				continue
			}
		}
	})

}

func (a *Agent) writeLoop(ctx context.Context) {
	conc.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-a.sendMQ:
				if atomic.LoadInt32(&a.state) != AgentStateRun {
					a.failMsg = data
					return
				}
				err := a.conn.Write(data)
				if err == nil {
					continue
				}
				a.failMsg = data
				a.OnError(err)
				return
			}
		}
	})
}

type Conn interface {
	Read(time.Duration) ([]byte, error)
	Write([]byte) error
	Close()
}
