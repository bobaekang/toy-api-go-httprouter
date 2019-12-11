package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Hello, World!")
}

func ArrestsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := []ArrestsAll{
		ArrestsAll{Year: 2017, Value: 1820},
		ArrestsAll{Year: 2018, Value: 1795},
	}
	writeOKResponse(w, data)
}

func ArrestsByOffenseClassPath(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := []ArrestsByOffenseClass{
		ArrestsByOffenseClass{Year: 2017, OffenseClass: 0, Value: 162},
		ArrestsByOffenseClass{Year: 2017, OffenseClass: 1, Value: 1277},
		ArrestsByOffenseClass{Year: 2017, OffenseClass: 2, Value: 81},
		ArrestsByOffenseClass{Year: 2018, OffenseClass: 0, Value: 421},
		ArrestsByOffenseClass{Year: 2018, OffenseClass: 1, Value: 1253},
		ArrestsByOffenseClass{Year: 2018, OffenseClass: 2, Value: 121},
	}
	writeOKResponse(w, data)
}

func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JsonResponse{Data: m}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&JsonErrorResponse{Error: &ApiError{Status: errorCode, Title: errorMsg}})
}
