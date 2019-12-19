package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/bobaekang/toy-api-go-httprouter/arrests"
)

func getIndex(s arrests.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Hello, World!")
	}
}

func getArrestsAll(s arrests.Service) func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data := s.GetArrestsAll()
		writeOKResponse(w, data)
	}
}

func getArrestsByOffenseClass(s arrests.Service) func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data := s.GetArrestsByOffenseClass()
		writeOKResponse(w, data)
	}
}

func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JSONResponse{Data: m}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&JSONErrorResponse{Error: &APIError{Status: errorCode, Title: errorMsg}})
}
