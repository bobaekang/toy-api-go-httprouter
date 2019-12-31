package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bobaekang/toy-api-go-httprouter/records"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	s := new(Storage)
	s.db = db

	return s
}

func (s *Storage) GetArrestsAll() (aa records.Records) {
	return fetchTableFromDB(s.db, "ArrestsAll")
}

func (s *Storage) GetArrestsByOffenseClass() (aa records.Records) {
	return fetchTableFromDB(s.db, "ArrestsByOffenseClass")
}

func fetchTableFromDB(db *sql.DB, table string) (aa records.Records) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		vv := make([]int, len(cols))
		vvPtrs := make([]interface{}, len(cols))

		for i := range vv {
			vvPtrs[i] = &vv[i]
		}

		if err = rows.Scan(vvPtrs...); err != nil {
			log.Fatal(err)
		}

		var groups []records.Group
		var value int

		for i, col := range cols {
			if col != "value" {
				groups = append(groups, records.Group{col, vv[i]})
			} else {
				value = vv[i]
			}
		}

		aa = append(aa, records.Record{groups, value})
	}

	return aa
}
