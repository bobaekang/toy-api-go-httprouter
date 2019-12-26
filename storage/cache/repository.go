package cache

import (
	"database/sql"
	"log"

	"github.com/bobaekang/toy-api-go-httprouter/arrests"
)

type Storage struct {
	db                    *sql.DB
	ArrestsAll            []arrests.All
	ArrestsByOffenseClass []arrests.ByOffenseClass
}

func NewStorage(db *sql.DB) *Storage {
	s := new(Storage)
	s.db = db
	s.ArrestsAll = FetchArrestsAll(db)
	s.ArrestsByOffenseClass = FetchArrestsByOffenseClass(db)

	return s
}

func (s *Storage) GetArrestsAll() []arrests.All {
	return s.ArrestsAll
}

func (s *Storage) GetArrestsByOffenseClass() []arrests.ByOffenseClass {
	return s.ArrestsByOffenseClass
}

func FetchArrestsAll(db *sql.DB) (aa []arrests.All) {
	rows, err := db.Query("SELECT * FROM ArrestsAll")
	defer rows.Close()

	for rows.Next() {
		a := new(arrests.All)

		if err = rows.Scan(&a.Year, &a.Value); err != nil {
			log.Fatal(err)
		}

		aa = append(aa, *a)
	}

	return aa
}

func FetchArrestsByOffenseClass(db *sql.DB) (aa []arrests.ByOffenseClass) {
	rows, err := db.Query("SELECT * FROM ArrestsByOffenseClass")
	defer rows.Close()

	for rows.Next() {
		a := new(arrests.ByOffenseClass)

		if err = rows.Scan(&a.Year, &a.OffenseClass, &a.Value); err != nil {
			log.Fatal(err)
		}

		aa = append(aa, *a)
	}

	return aa
}
