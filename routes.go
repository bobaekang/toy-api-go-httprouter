package main

import "github.com/julienschmidt/httprouter"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index,
		},
		Route{
			"ArrestsIndex",
			"GET",
			"/arrests",
			ArrestsIndex,
		},
		Route{
			"ArrestsByOffenseClassPath",
			"GET",
			"/arrests/by-offense-class",
			ArrestsByOffenseClassPath,
		},
	}

	return routes
}
