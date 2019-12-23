package sqlite

import (
	"database/sql"
	"log"

	"github.com/bobaekang/toy-api-go-httprouter/arrests"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	s := new(Storage)
	s.db = db

	return s
}

func (s *Storage) GetArrestsAll() (aa []arrests.All) {
	rows, err := s.db.Query("SELECT * FROM ArrestsAll")
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

func (s *Storage) GetArrestsByOffenseClass() (aa []arrests.ByOffenseClass) {
	rows, err := s.db.Query("SELECT * FROM ArrestsByOffenseClass")
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
