package main

import (
	"net/http"
	"regexp"
)

type Handler func(http.ResponseWriter, *http.Request, []string)

type route struct {
	pattern *regexp.Regexp
	handler Handler
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if true == route.pattern.MatchString(r.URL.Path) {
			matches := route.pattern.FindStringSubmatch(r.URL.Path)
			route.handler(w, r, matches)

			return
		}
	}

	http.NotFound(w, r)
}
