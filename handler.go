package main

import (
	"fmt"
	"net/http"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request, matches []string) {
	app, exists := getApp(matches)

	if false == exists || nil == app {
		http.NotFound(w, r)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		_, _ = fmt.Fprintf(w, Response{Ok: true, Errors: []string{"Error parse body"}}.ToJson())

		return
	}

	_, _ = fmt.Fprintf(w, Response{Ok: true}.ToJson())

	command := r.Form["command"]

	if len(command) > 0 {
		go runCommands(command, app)
	}
}

func StatRouterHandler(w http.ResponseWriter, r *http.Request, matches []string) {
	app, exists := getApp(matches)

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

	_, _ = fmt.Fprintf(w, Response{Ok: true, Data: data}.ToJson())
}

func getApp(matches []string) (*App, bool) {
	if len(matches) < 2 {
		return nil, false
	}

	config, exists := apps[matches[1]]
	if false == exists {
		return nil, false
	}

	return config, true
}
