package rest

import (
	"regexp"
	"strings"

	"github.com/bobaekang/toy-api-go-httprouter/data"
	"github.com/julienschmidt/httprouter"
)

func toPath(tableName string) string {
	re := regexp.MustCompile("([A-Z])")
	path := re.ReplaceAllString(tableName, "-$1")
	path = strings.Replace(path, "-", "/", 1)
	path = strings.Replace(path, "-By", "/By", 1)
	path = strings.Replace(path, "Ref-", "Ref/", 1)

	return strings.ToLower(path)
}

func NewRouter(s data.Service) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", logger(getIndex(s)))
	router.GET(toPath("Arrests"), logger(getTable(s, "Arrests")))
	router.GET(toPath("ArrestsByOffenseClass"), logger(getTable(s, "ArrestsByOffenseClass")))
	router.GET(toPath("RefOffenseClass"), logger(getRefTable(s, "RefOffenseClass")))

	return router
}
