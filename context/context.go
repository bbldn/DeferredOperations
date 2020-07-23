package context

import Config "deferredOperations/config"
import Application "deferredOperations/application"

type Context struct {
	Config Config.Config
	Apps   map[string]*Application.App
}
