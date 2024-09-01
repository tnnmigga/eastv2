package main

import (
	"eastv2/game/modules/play"

	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(play.New())
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
