package user

import (
	"eastv2/game/modules/play/user/model"

	"github.com/tnnmigga/corev2/iface"
)

var manager *Manager

type Manager struct {
	iface.IModule
	cache   map[uint64]*model.Model
	waiting map[uint64][]func(*model.Model, error)
}

func Init(m iface.IModule) {
	manager = &Manager{
		IModule: m,
		cache:   map[uint64]*model.Model{},
		waiting: map[uint64][]func(*model.Model, error){},
	}
}
