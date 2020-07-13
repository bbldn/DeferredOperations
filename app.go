package main

import (
	"github.com/bigkevmcd/go-configparser"
	"os/exec"
)

type App struct {
	Processes map[string]*exec.Cmd
	Commands  map[int]string
	Config    configparser.Dict
}

func (c *App) Load(config configparser.Dict) {
	c.Processes = make(map[string]*exec.Cmd)
	c.Commands = make(map[int]string)
	c.Config = config
}
