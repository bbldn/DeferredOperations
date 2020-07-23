package router

import (
	"net/http"
	"regexp"
	"strings"
)

type action func(http.ResponseWriter, *http.Request, []string)

type route struct {
	pattern *regexp.Regexp
	action  action
}

type Router struct {
	routes []*route
}

func (h *Router) AddAction(pattern *regexp.Regexp, action action) {
	h.routes = append(h.routes, &route{pattern, action})
}

func (h Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimRight(r.URL.Path, "/")
	for _, route := range h.routes {
		if true == route.pattern.MatchString(path) {
			matches := route.pattern.FindStringSubmatch(path)
			route.action(w, r, matches)

			return
		}
	}

	http.NotFound(w, r)
}
