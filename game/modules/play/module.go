package play

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/impl"
	"eastv2/game/modules/play/userops"

	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

type play struct {
	*domain.Domain
}

func New() iface.IModule {
	m := &play{
		Domain: domain.New(basic.NewEventLoop(100000)),
	}
	userops.Init(m)
	impl.Init(m.Domain)
	return m
}

func (m *play) Name() string {
	return "play"
}
