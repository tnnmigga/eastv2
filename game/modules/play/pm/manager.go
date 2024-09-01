package pm

import (
	"github.com/tnnmigga/corev2/iface"
)

var manager *Manager

type Manager struct {
	iface.IModule
	cache   map[uint64]*UserData
	waiting map[uint64][]func(*UserData, error)
}

func Init(m iface.IModule) {
	manager = &Manager{
		IModule: m,
		cache:   map[uint64]*UserData{},
		waiting: map[uint64][]func(*UserData, error){},
	}
}

func HandleUserMsg[T any](fn func(user *UserData, req *T)) {

}
