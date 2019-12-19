package main

import (
	"log"
	"net/http"

	"github.com/bobaekang/toy-api-go-httprouter/http/rest"
)

func main() {
	router := rest.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
