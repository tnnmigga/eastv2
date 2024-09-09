package account

import (
	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/module"
)

type account struct {
	iface.IModule
}

func New() iface.IModule {
	m := &account{
		IModule: module.NewConcurrency("account"),
	}
	m.initHandler()
	return m
}
