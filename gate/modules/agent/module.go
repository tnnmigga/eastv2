package agent

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/log"
)

const (
	AgentTypeTCP       = "tcp"
	AgentTypeWebSocket = "websocket"
)

type agent struct {
	iface.IModule
	agentType string
	listener  IListener
	manager   *AgentManager
}

type IListener interface {
	Run() error
	Close()
}

func New(agentType string) iface.IModule {
	m := basic.NewConcurrency("agent")
	a := &agent{
		IModule:   m,
		agentType: agentType,
		manager:   NewAgentManager(m),
	}
	switch agentType {
	case AgentTypeTCP:
		a.listener = NewTCPListener(a.manager)
	case AgentTypeWebSocket:
		a.listener = NewWebSocketListener(a.manager)
	default:
		log.Panicf("unknown agent type: %s", agentType)
	}
	a.initHandler()
	return a
}

func (m *agent) Run() error {
	return m.listener.Run()
}

func (m *agent) Exit() error {
	m.listener.Close()
	for _, agent := range m.manager.agents {
		agent.cancel()
		agent.conn.Close()
	}
	return nil
}
