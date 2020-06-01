package main

import (
	"deferredOperations/config"
	"deferredOperations/console"
	"deferredOperations/response"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var appConfig config.Config
var Processes map[string]*exec.Cmd
var Commands map[int]string

func runCommand(command string) {
	commands := strings.Split(command, " ")
	commands = append([]string{appConfig.Values["APP_PATH"]}, commands...)

	cmd := exec.Command("php", commands...)
	err := cmd.Start()
	if err == nil {
		key := fmt.Sprintf("%d %s", cmd.Process.Pid, command)
		Processes[key] = cmd
		_ = cmd.Wait()
		delete(Processes, key)
	}
}

func runCommands(commands []string) {
	startIndex := len(Commands)

	for key, command := range commands {
		Commands[startIndex+key+1] = command
	}

	for key, command := range commands {
		runCommand(command)
		delete(Commands, key+startIndex+1)
	}
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		_, _ = fmt.Fprintf(w, response.Response{Ok: true, Errors: []string{"Error parse body"}}.ToJson())

		return
	}

	_, _ = fmt.Fprintf(w, response.Response{Ok: true}.ToJson())

	command := r.Form["command"]

	if len(command) > 0 {
		go runCommands(command)
	}
}

func StatRouterHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys := make([]string, len(Processes))
	i := 0
	for p := range Processes {
		keys[i] = p
		i++
	}

	i = 0
	commands := make([]string, len(Commands))
	for _, command := range Commands {
		commands[i] = command
		i++
	}

	data := make(map[string]interface{})
	data["Processes"] = keys
	data["Commands"] = commands

	_, _ = fmt.Fprintf(w, response.Response{Ok: true, Data: data}.ToJson())
}

func main() {
	values, err := console.ParseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = appConfig.Load(values)
	if err != nil {
		log.Fatal(err)
		return
	}

	Processes = make(map[string]*exec.Cmd)
	Commands = make(map[int]string)

	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/stat", StatRouterHandler)

	addr := fmt.Sprintf("%s:%s", appConfig.Values["ADDRESS"], appConfig.Values["PORT"])
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error start server:", err)
	}
}
