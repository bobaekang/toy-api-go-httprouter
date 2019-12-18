package main

import "github.com/julienschmidt/httprouter"

func NewRouter(routes []Route) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle
		handle = Logger(route.HandlerFunc)

		router.Handle(route.Method, route.Path, handle)
	}

	return router
}
