package main

import (
	"monitoring/internal/app"
	"monitoring/internal/ui"
)

func main() {
	ui.ShowIntroduction()
	ui.ShowMenu()

	command := app.ReadCommand()
	app.ExecuteCommand(command)
}
