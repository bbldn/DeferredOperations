package controllers

import (
	Helpers "deferredOperations/helpers"
	Application "deferredOperations/helpers/application"
	Context "deferredOperations/helpers/context"
	"fmt"
	"net/http"
)

type Main struct {
	Context Context.Context
}

func (m Main) HomeAction(w http.ResponseWriter, r *http.Request, matches []string) {
	app, exists := m.getApp(matches)

	if false == exists || nil == app {
		http.NotFound(w, r)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		_, _ = fmt.Fprintf(w, Helpers.Response{Ok: true, Errors: []string{"Error parse body"}}.ToJson())

		return
	}

	_, _ = fmt.Fprintf(w, Helpers.Response{Ok: true}.ToJson())

	command := r.Form["command"]

	if len(command) > 0 {
		go Helpers.RunCommands(command, app)
	}
}

func (m Main) StatAction(w http.ResponseWriter, r *http.Request, matches []string) {
	app, exists := m.getApp(matches)

	if false == exists || nil == app {
		http.NotFound(w, r)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	keys := make([]string, len(app.Processes))
	i := 0
	for p := range app.Processes {
		keys[i] = p
		i++
	}

	i = 0
	commands := make([]string, len(app.Commands))
	for _, command := range app.Commands {
		commands[i] = command
		i++
	}

	data := make(map[string]interface{})
	data["Processes"] = keys
	data["Commands"] = commands

	_, _ = fmt.Fprintf(w, Helpers.Response{Ok: true, Data: data}.ToJson())
}

func (m Main) getApp(matches []string) (*Application.App, bool) {
	if len(matches) < 2 {
		return nil, false
	}

	config, exists := m.Context.Apps[matches[1]]
	if false == exists {
		return nil, false
	}

	return config, true
}
