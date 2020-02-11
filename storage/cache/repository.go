package cache

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bobaekang/toy-api-go-httprouter/data"
)

type Storage struct {
	db        *sql.DB
	tables    map[string]data.Table
	refTables map[string]data.RefTable
}

func NewStorage(db *sql.DB) *Storage {
	mt := make(map[string]data.Table)
	mr := make(map[string]data.RefTable)
	for _, tableName := range fetchTableNamesFromDB(db) {
		if tableName[:3] == "Ref" {
			mr[tableName] = fetchRefTableFromDB(db, tableName)
		} else {
			mt[tableName] = fetchTableFromDB(db, tableName)
		}
	}

	s := new(Storage)
	s.db = db
	s.tables = mt
	s.refTables = mr

	return s
}

func (s *Storage) GetTable(tableName string) data.Table {
	return s.tables[tableName]
}

func (s *Storage) GetRefTable(tableName string) data.RefTable {
	return s.refTables[tableName]
}

func fetchTableNamesFromDB(db *sql.DB) []string {
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		tableNames = append(tableNames, name)
	}

	return tableNames
}

func fetchTableFromDB(db *sql.DB, tableName string) (table data.Table) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
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

		row := make(data.Row, len(cols))
		for i, col := range cols {
			row[i] = data.Variable{col, values[i]}
		}

		table = append(table, row)
	}

	return table
}

func fetchRefTableFromDB(db *sql.DB, tableName string) (refTable data.RefTable) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var value string
	for rows.Next() {
		if err = rows.Scan(&id, &value); err != nil {
			log.Fatal(err)
		}

		refTable = append(refTable, data.RefRow{id, value})
	}

	return refTable
}
