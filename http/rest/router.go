package rest

import (
	"github.com/bobaekang/toy-api-go-httprouter/data"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(s data.Service) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", logger(getIndex(s)))
	router.GET("/arrests", logger(getArrestsAll(s)))
	router.GET("/arrests/by-offense-class", logger(getArrestsByOffenseClass(s)))

	return router
}
