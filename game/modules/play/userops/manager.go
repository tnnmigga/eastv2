package userops

import (
	"eastv2/game/modules/play/userops/muser"

	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/module"
)

var manager *Manager

type Manager struct {
	iface.IModule
	cache   map[uint64]*muser.Model
	waiting map[uint64][]func(*muser.Model, error)
}

func Init(m iface.IModule) {
	manager = &Manager{
		IModule: m,
		cache:   map[uint64]*muser.Model{},
		waiting: map[uint64][]func(*muser.Model, error){},
	}
	module.Handle(m, onC2SPackage)
}
