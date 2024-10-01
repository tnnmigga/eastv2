package main

import (
	"eastv2/door/modules/account"
	"eastv2/door/modules/web"


	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(account.New(), web.New())
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
