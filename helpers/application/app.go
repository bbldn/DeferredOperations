package application

import (
	GoConfigParser "github.com/bigkevmcd/go-configparser"
	"os/exec"
)

type App struct {
	Processes map[string]*exec.Cmd
	Commands  map[int]string
	Config    GoConfigParser.Dict
}

func (c *App) Load(config GoConfigParser.Dict) {
	c.Processes = make(map[string]*exec.Cmd)
	c.Commands = make(map[int]string)
	c.Config = config
}
