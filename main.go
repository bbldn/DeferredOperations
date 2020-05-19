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
var processes map[string]*exec.Cmd

func runCommand(command string) {
	commands := strings.Split(command, " ")
	commands = append([]string{appConfig.Values["APP_PATH"]}, commands...)

	cmd := exec.Command("php", commands...)
	err := cmd.Start()
	if err == nil {
		key := fmt.Sprintf("%d %s", cmd.Process.Pid, command)
		processes[key] = cmd
		_ = cmd.Wait()
		delete(processes, key)
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

	command := r.Form.Get("command")

	if len(command) > 0 {
		go runCommand(command)
	}
}

func StatRouterHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys := make([]string, len(processes))
	i := 0
	for p := range processes {
		keys[i] = p
		i++
	}
	data := make(map[string]interface{})
	data["processes"] = keys

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

	processes = make(map[string]*exec.Cmd)

	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/stat", StatRouterHandler)

	addr := fmt.Sprintf("%s:%s", appConfig.Values["ADDRESS"], appConfig.Values["PORT"])
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error start server:", err)
	}
}
