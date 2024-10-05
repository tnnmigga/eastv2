package agent

import (
	"eastv2/pb"

	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/message"
)

func (m *agent) register() {
	message.Handle(m, m.onS2CPackage)
}

func (m *agent) onS2CPackage(pkg *pb.S2CMsg) {
	agent := m.manager.GetAgent(pkg.UserID)
	if agent == nil {
		log.Warnf("agent not found %d", pkg.UserID)
		return
	}
	select {
	case agent.sendMQ <- pkg.Body:
	default:
		log.Errorf("agent send mq full! %d", pkg.UserID)
	}
}
