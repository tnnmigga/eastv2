package web

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/utils/hs"
)

type web struct {
	iface.IModule
	*hs.HttpService
}

func New() iface.IModule {
	m := &web{
		IModule: basic.NewConcurrency("account"),
	}
	m.initHandle()
	return m
}
