package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bobaekang/toy-api-go-httprouter/http/rest"
	"github.com/bobaekang/toy-api-go-httprouter/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn := sqliteConnection("./database.db")
	defer conn.Close()

	s := sqlite.NewStorage(conn)
	router := rest.NewRouter(s)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func sqliteConnection(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("note: connected to SQLite database.")

	return db
}
