package main

import (
	"deferredOperations/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

var appConfig config.Config
var apps map[string]*App

func main() {
	values, err := ParseArgs(os.Args)
	if err != nil {
		log.Fatal(err)

		return
	}

	err = appConfig.Load(values)
	if err != nil {
		log.Fatal(err)

		return
	}

	apps = make(map[string]*App)

	for _, section := range appConfig.Values.Sections() {
		if "DEFAULTS" != section {
			app := App{}
			config, _ := appConfig.Values.Items(section)
			app.Load(config)
			apps[section] = &app
		}
	}

	server := RegexpHandler{}
	server.HandleFunc(regexp.MustCompile(`^/(.+)$`), HomeRouterHandler)
	server.HandleFunc(regexp.MustCompile(`^/(.+)?/stat/?$`), StatRouterHandler)

	address, _ := appConfig.Values.Get("DEFAULTS", "ADDRESS")
	port, _ := appConfig.Values.Get("DEFAULTS", "PORT")

	addr := fmt.Sprintf("%s:%s", address, port)
	err = http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatal("Error start server:", err)
	}
}
