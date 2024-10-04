package account

import (
	"fmt"

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

func tokenKey(token string) string {
	return fmt.Sprintf("token:{%s}", token)
}
