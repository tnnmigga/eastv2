package web

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/conf"
	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/utils/hs"
)

type web struct {
	iface.IModule
	*hs.HttpService
}

func New() iface.IModule {
	m := &web{
		IModule:     basic.NewConcurrency(),
		HttpService: hs.NewHttpService(),
	}
	m.initHandle()
	m.ListenAndServe(conf.String("web.addr", ""))
	return m
}

func (m *web) Name() string {
	return "web"
}

func (m *web) Exit() error {
	return m.Stop()
}
