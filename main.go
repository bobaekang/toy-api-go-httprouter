package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/bobaekang/toy-api-go-httprouter/data"
	"github.com/bobaekang/toy-api-go-httprouter/http/rest"
	"github.com/bobaekang/toy-api-go-httprouter/storage/memory"
	"github.com/bobaekang/toy-api-go-httprouter/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	storageType := flag.String("storage", "sqlite", "storage type [memory, sqlite]")
	flag.Parse()

	var dataService data.Service

	switch *storageType {
	case "memory":
		s := memory.NewStorage()
		dataService = data.NewService(s)

	case "sqlite":
		conn := sqliteConnection("./database.db")
		s := sqlite.NewStorage(conn)
		conn.Close()

		dataService = data.NewService(s)
	}

	router := rest.NewRouter(dataService)
	log.Println("note: listing on port 8080.")
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
