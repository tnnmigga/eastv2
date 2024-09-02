package play

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/impl"
	"eastv2/game/modules/play/user"

	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/module"
)

type play struct {
	iface.IModule
	*domain.Domain
}

func New() iface.IModule {
	m := &play{
		IModule: module.NewEventLoop("play", 100000),
	}
	m.Domain = domain.New(m)
	user.Init(m)
	impl.Init(m.Domain)
	return m
}
