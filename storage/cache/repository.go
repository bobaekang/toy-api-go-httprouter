package cache

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bobaekang/toy-api-go-httprouter/records"
)

type Storage struct {
	db                    *sql.DB
	ArrestsAll            records.Records
	ArrestsByOffenseClass records.Records
}

func NewStorage(db *sql.DB) *Storage {
	s := new(Storage)
	s.db = db
	s.ArrestsAll = fetchTableFromDB(db, "ArrestsAll")
	s.ArrestsByOffenseClass = fetchTableFromDB(db, "ArrestsByOffenseClass")

	return s
}

func (s *Storage) GetArrestsAll() records.Records {
	return s.ArrestsAll
}

func (s *Storage) GetArrestsByOffenseClass() records.Records {
	return s.ArrestsByOffenseClass
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
