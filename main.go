package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/bobaekang/toy-api-go-httprouter/http/rest"
	"github.com/bobaekang/toy-api-go-httprouter/records"
	"github.com/bobaekang/toy-api-go-httprouter/storage/cache"
	"github.com/bobaekang/toy-api-go-httprouter/storage/memory"
	"github.com/bobaekang/toy-api-go-httprouter/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	storageType := flag.String("storage", "sqlite", "storage type [cache, memory, sqlite]")
	flag.Parse()

	var recordsService records.Service

	switch *storageType {
	case "cache":
		conn := sqliteConnection("./database.db")
		defer conn.Close()

		s := cache.NewStorage(conn)
		recordsService = records.NewService(s)

	case "memory":
		s := memory.NewStorage()
		recordsService = records.NewService(s)

	case "sqlite":
		conn := sqliteConnection("./database.db")
		defer conn.Close()

		s := sqlite.NewStorage(conn)
		recordsService = records.NewService(s)
	}

	router := rest.NewRouter(recordsService)
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
