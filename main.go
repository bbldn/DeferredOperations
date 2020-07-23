package main

import (
	Application "deferredOperations/application"
	Context "deferredOperations/context"
	Controllers "deferredOperations/controllers"
	Router "deferredOperations/helpers/router"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var context Context.Context

func init() {
	err := context.Config.Load()
	if err != nil {
		log.Fatal(err)

		return
	}

	context.Apps = make(map[string]*Application.App)
	for _, section := range context.Config.Values.Sections() {
		if "DEFAULTS" != section {
			app := Application.App{}
			c, _ := context.Config.Values.Items(section)
			app.Load(c)
			context.Apps[section] = &app
		}
	}
}

func main() {
	mainController := Controllers.Main{Context: context}

	server := Router.Router{}
	server.AddAction(regexp.MustCompile(`^/(.+)?/stat$`), mainController.StatAction)
	server.AddAction(regexp.MustCompile(`^/([^/]+)$`), mainController.HomeAction)

	//errors ignored because config have been validated
	address, _ := context.Config.Values.Get("DEFAULTS", "ADDRESS")
	port, _ := context.Config.Values.Get("DEFAULTS", "PORT")

	addr := fmt.Sprintf("%s:%s", address, port)
	err := http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatal("Error start router:", err)
	}
}
