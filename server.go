package main

import (
	"net/http"
	"regexp"
	"strings"
)

type Handler func(http.ResponseWriter, *http.Request, []string)

type route struct {
	pattern *regexp.Regexp
	handler Handler
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimRight(r.URL.Path, "/")

	for _, route := range h.routes {
		if true == route.pattern.MatchString(path) {
			matches := route.pattern.FindStringSubmatch(path)
			route.handler(w, r, matches)

			return
		}
	}

	http.NotFound(w, r)
}
