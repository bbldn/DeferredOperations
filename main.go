package main

import (
	. "deferredOperations/config"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var config Config
var apps map[string]*App

func init() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)

		return
	}

	apps = make(map[string]*App)

	for _, section := range config.Values.Sections() {
		if "DEFAULTS" != section {
			app := App{}
			c, _ := config.Values.Items(section)
			app.Load(c)
			apps[section] = &app
		}
	}
}

func main() {
	server := RegexpHandler{}
	server.HandleFunc(regexp.MustCompile(`^/(.+)?/stat$`), StatRouterHandler)
	server.HandleFunc(regexp.MustCompile(`^/([^/]+)$`), HomeRouterHandler)

	address, _ := config.Values.Get("DEFAULTS", "ADDRESS")
	port, _ := config.Values.Get("DEFAULTS", "PORT")

	addr := fmt.Sprintf("%s:%s", address, port)
	err := http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatal("Error start server:", err)
	}
}
