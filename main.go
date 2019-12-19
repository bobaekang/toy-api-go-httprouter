package main

import (
	"log"
	"net/http"
)

func main() {
	routes := AllRoutes()
	router := NewRouter(routes...)
	log.Fatal(http.ListenAndServe(":8080", router))
}
