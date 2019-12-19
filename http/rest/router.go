package rest

import (
	"github.com/bobaekang/toy-api-go-httprouter/arrests"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(s arrests.Service) *httprouter.Router {
	router := httprouter.New()

	handleIndex := logger(getIndex(s))
	handleArrestsAll := logger(getArrestsAll(s))
	handleArrestsByOffenseClass := logger(getArrestsByOffenseClass(s))

	router.GET("/", handleIndex)
	router.GET("/arrests", handleArrestsAll)
	router.GET("/arrests/by-offense-class", handleArrestsByOffenseClass)

	return router
}
