package main

import (
	"log"
	"net/http"

	"github.com/bobaekang/toy-api-go-httprouter/http/rest"
	"github.com/bobaekang/toy-api-go-httprouter/storage/memory"
)

func main() {
	s := memory.NewStorage()
	router := rest.NewRouter(s)
	log.Fatal(http.ListenAndServe(":8080", router))
}
