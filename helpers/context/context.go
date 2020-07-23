package context

import Config "deferredOperations/helpers/config"
import Application "deferredOperations/helpers/application"

type Context struct {
	Config Config.Config
	Apps   map[string]*Application.App
}
