package main

import (
	. "github.com/bigkevmcd/go-configparser"
	"os/exec"
)

type App struct {
	Processes map[string]*exec.Cmd
	Commands  map[int]string
	Config    Dict
}

func (c *App) Load(config Dict) {
	c.Processes = make(map[string]*exec.Cmd)
	c.Commands = make(map[int]string)
	c.Config = config
}
