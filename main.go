package main

import (
	"log"
	"net/http"

	"github.com/bobaekang/toy-api-go-httprouter/http/rest"
)

func main() {
	routes := rest.AllRoutes()
	router := rest.NewRouter(routes...)
	log.Fatal(http.ListenAndServe(":8080", router))
}
