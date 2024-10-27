package play

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/impl"
	"eastv2/game/modules/play/userops"

	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

const moduleName = "play"

type play struct {
	*domain.Domain
}

func New() iface.IModule {
	m := &play{
		Domain: domain.New(basic.NewEventLoop(moduleName, basic.DefaultMQLen)),
	}
	userops.Init(m)
	impl.Init(m.Domain)
	return m
}
