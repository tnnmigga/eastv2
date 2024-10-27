package account

import (
	"fmt"

	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

const moduleName = "account"

type account struct {
	iface.IModule
}

func New() iface.IModule {
	m := &account{
		IModule: basic.NewConcurrency(moduleName),
	}
	m.register()
	return m
}

func tokenKey(token string) string {
	return fmt.Sprintf("token:{%s}", token)
}
