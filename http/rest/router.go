package rest

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	routes := allRoutes()
	for _, r := range routes {
		var handle httprouter.Handle
		handle = logger(r.HandlerFunc)

		router.Handle(r.Method, r.Path, handle)
	}

	return router
}
