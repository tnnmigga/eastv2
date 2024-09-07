package main

import (
	"eastv2/gate/modules/agent"

	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(agent.New(agent.AgentTypeTCP))
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
