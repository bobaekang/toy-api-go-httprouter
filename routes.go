package main

import "github.com/julienschmidt/httprouter"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

func AllRoutes() []Route {
	return []Route{
		{
			"Index",
			"GET",
			"/",
			Index,
		},
		{
			"ArrestsIndex",
			"GET",
			"/arrests",
			ArrestsIndex,
		},
		{
			"ArrestsByOffenseClassPath",
			"GET",
			"/arrests/by-offense-class",
			ArrestsByOffenseClassPath,
		},
	}
}
