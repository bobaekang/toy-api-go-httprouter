package rest

import "github.com/julienschmidt/httprouter"

func NewRouter(routes ...Route) *httprouter.Router {
	router := httprouter.New()
	for _, r := range routes {
		var handle httprouter.Handle
		handle = Logger(r.HandlerFunc)

		router.Handle(r.Method, r.Path, handle)
	}

	return router
}
