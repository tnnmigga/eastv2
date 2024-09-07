package play

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/impl"
	"eastv2/game/modules/play/userops"

	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/module"
)

type play struct {
	*domain.Domain
}

func New() iface.IModule {
	m := &play{
		Domain: domain.New(module.NewEventLoop("play", 100000)),
	}
	userops.Init(m)
	impl.Init(m.Domain)
	return m
}
