package userops

import (
	"eastv2/game/modules/play/userops/userdata"

	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/message"
)

var manager *Manager

type Manager struct {
	iface.IModule
	cache   map[uint64]*userdata.Entity
	waiting map[uint64][]func(*userdata.Entity, error)
}

func Init(m iface.IModule) {
	manager = &Manager{
		IModule: m,
		cache:   map[uint64]*userdata.Entity{},
		waiting: map[uint64][]func(*userdata.Entity, error){},
	}
	message.Handle(m, onC2SPackage)
}
