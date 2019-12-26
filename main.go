package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bobaekang/toy-api-go-httprouter/arrests"
	"github.com/bobaekang/toy-api-go-httprouter/http/rest"
	"github.com/bobaekang/toy-api-go-httprouter/storage/memory"
	"github.com/bobaekang/toy-api-go-httprouter/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

// StorageType defines available storage types
type StorageType int

const (
	// Memory will store data in memory
	Memory StorageType = iota
	// Sqlite will store data in SQLite .db file saved on disk
	Sqlite
)

func main() {
	storageType := Sqlite

	var arrestsService arrests.Service

	switch storageType {
	case Memory:
		s := memory.NewStorage()
		arrestsService = arrests.NewService(s)

	case Sqlite:
		conn := sqliteConnection("./database.db")
		defer conn.Close()

		s := sqlite.NewStorage(conn)
		arrestsService = arrests.NewService(s)
	}

	router := rest.NewRouter(arrestsService)
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
