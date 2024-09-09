package main

import (
	"eastv2/door/modules/account"

	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(account.New())
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
