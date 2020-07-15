package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func runCommand(command string, app *App) {
	handler, exists := app.Config["HANDLER"]

	var commands []string
	if true == exists {
		commands = strings.Split(command, " ")
		commands = append([]string{app.Config["APP_PATH"]}, commands...)
	} else {
		commands = strings.Split(command, " ")
		handler = app.Config["APP_PATH"]
	}

	cmd := exec.Command(handler, commands...)
	err := cmd.Start()
	if err == nil {
		key := fmt.Sprintf("%d %s", cmd.Process.Pid, command)
		app.Processes[key] = cmd
		_ = cmd.Wait()
		delete(app.Processes, key)
	}
}

func runCommands(commands []string, app *App) {
	startIndex := len(app.Commands)

	for key, command := range commands {
		app.Commands[startIndex+key+1] = command
	}

	for key, command := range commands {
		runCommand(command, app)
		delete(app.Commands, key+startIndex+1)
	}
}
