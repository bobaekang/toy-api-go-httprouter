package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/arrests", ArrestsIndex)
	router.GET("/arrests/by-offense-class", ArrestsByOffenseClassPath)

	log.Fatal(http.ListenAndServe(":8080", router))
}
