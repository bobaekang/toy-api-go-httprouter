package main

import (
	"encoding/json"
	"fmt"
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

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Hello, World!")
}

func ArrestsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Arrests all")
	data := ArrestsAll{
		ArrestsAllOne{Year: 2017, Value: 1820},
		ArrestsAllOne{Year: 2018, Value: 1795},
	}
	json.NewEncoder(w).Encode(data)
}

func ArrestsByOffenseClassPath(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Arrests by offense class")
	data := ArrestsByOffenseClass{
		ArrestsByOffenseClassOne{Year: 2017, OffenseClass: 0, Value: 162},
		ArrestsByOffenseClassOne{Year: 2017, OffenseClass: 1, Value: 1277},
		ArrestsByOffenseClassOne{Year: 2017, OffenseClass: 2, Value: 81},
		ArrestsByOffenseClassOne{Year: 2018, OffenseClass: 0, Value: 421},
		ArrestsByOffenseClassOne{Year: 2018, OffenseClass: 1, Value: 1253},
		ArrestsByOffenseClassOne{Year: 2018, OffenseClass: 2, Value: 121},
	}
	json.NewEncoder(w).Encode(data)
}
