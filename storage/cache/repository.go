package cache

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bobaekang/toy-api-go-httprouter/data"
)

type Storage struct {
	db     *sql.DB
	tables map[string]data.Table
}

func NewStorage(db *sql.DB) *Storage {
	m := make(map[string]data.Table)
	m["ArrestsAll"] = fetchTableFromDB(db, "ArrestsAll")
	m["ArrestsByOffenseClass"] = fetchTableFromDB(db, "ArrestsByOffenseClass")

	s := new(Storage)
	s.db = db
	s.tables = m

	return s
}

func (s *Storage) GetArrestsAll() data.Table {
	return s.tables["ArrestsAll"]
}

func (s *Storage) GetArrestsByOffenseClass() data.Table {
	return s.tables["ArrestsByOffenseClass"]
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
