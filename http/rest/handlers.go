package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bobaekang/toy-api-go-httprouter/data"
	"github.com/julienschmidt/httprouter"
)

func getIndex(s data.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Hello, World!")
	}
}

func getTable(s data.Service, tableName string) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data := s.GetTable(tableName)
		query := r.URL.Query()
		queryResult := runQuery(data, query)
		writeOKResponse(w, queryResult)
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

func runQuery(table data.Table, query map[string][]string) data.Table {
	if yearMin := query["yearMin"]; len(yearMin) > 0 {
		value, err := strconv.Atoi(yearMin[0])
		if err != nil {
			return nil
		}
		table = table.Filter("year", ">=", value)
	}

	if yearMax := query["yearMax"]; len(yearMax) > 0 {
		value, err := strconv.Atoi(yearMax[0])
		if err != nil {
			return nil
		}
		table = table.Filter("year", "<=", value)
	}

	if sortBy := query["sortBy"]; len(sortBy) > 0 {
		for _, s := range strings.Split(sortBy[0], " ") {
			byOrder := strings.Split(s, ":")
			by := byOrder[0]
			order := "asc"

			if len(byOrder) > 1 {
				order = strings.ToLower(byOrder[1])
				if order != "asc" && order != "desc" {
					return nil
				}
			}

			table = table.SortBy(by, order)
		}
	}

	if fields := query["variables"]; len(fields) > 0 {
		table = table.Select(fields...)
	}

	return table
}
