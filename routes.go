package main

import "github.com/julienschmidt/httprouter"

type Route struct {
	Method string
	Path   string
	Handle httprouter.Handle
}

type Routes []Route

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Handle)
	}

	return router
}

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/arrests",
		ArrestsIndex,
	},
	Route{
		"GET",
		"/arrests/by-offense-class",
		ArrestsByOffenseClassPath,
	},
}
