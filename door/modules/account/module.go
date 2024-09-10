package account

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

type account struct {
	iface.IModule
}

func New() iface.IModule {
	m := &account{
		IModule: basic.NewConcurrency("account"),
	}
	m.initHandler()
	return m
}
