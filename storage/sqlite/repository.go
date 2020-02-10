package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bobaekang/toy-api-go-httprouter/data"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	s := new(Storage)
	s.db = db

	return s
}

func (s *Storage) GetTable(tableName string) data.Table {
	return fetchTableFromDB(s.db, tableName)
}

func fetchTableFromDB(db *sql.DB, tableName string) (table data.Table) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
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
		values := make([]int, len(cols))
		valuesPtrs := make([]interface{}, len(cols))

		for i := range values {
			valuesPtrs[i] = &values[i]
		}

		if err = rows.Scan(valuesPtrs...); err != nil {
			log.Fatal(err)
		}

		var variables []data.Variable
		var value int

		for i, col := range cols {
			if col != "value" {
				variables = append(variables, data.Variable{col, values[i]})
			} else {
				value = values[i]
			}
		}

		table = append(table, data.Row{variables, value})
	}

	return table
}
