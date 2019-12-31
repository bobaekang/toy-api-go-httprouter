package rest

import (
	"github.com/bobaekang/toy-api-go-httprouter/records"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(s records.Service) *httprouter.Router {
	router := httprouter.New()

	handleIndex := logger(getIndex(s))
	handleArrestsAll := logger(getArrestsAll(s))
	handleArrestsByOffenseClass := logger(getArrestsByOffenseClass(s))

	router.GET("/", handleIndex)
	router.GET("/arrests", handleArrestsAll)
	router.GET("/arrests/by-offense-class", handleArrestsByOffenseClass)

	return router
}
