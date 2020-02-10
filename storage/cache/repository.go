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

func (s *Storage) GetTable(tableName string) data.Table {
	return s.tables[tableName]
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

		var row data.Row

		for i, col := range cols {
			row = append(row, data.Variable{col, values[i]})
		}

		table = append(table, row)
	}

	return table
}
