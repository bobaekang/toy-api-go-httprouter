package main

import "github.com/julienschmidt/httprouter"

type Route struct {
	Method string
	Path   string
	Handle httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
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

	return routes
}
