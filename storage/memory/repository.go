package memory

import (
	"github.com/bobaekang/toy-api-go-httprouter/data"
)

type Storage struct {
	tables    map[string]data.Table
	refTables map[string]data.RefTable
}

func NewStorage() *Storage {
	mt := make(map[string]data.Table)
	mt["Arrests"] = data.Table{
		{{"year", 2017}, {"value", 1820}},
		{{"year", 2018}, {"value", 1795}},
	}
	mt["ArrestsByOffenseClass"] = data.Table{
		{{"year", 2017}, {"offenseclass", 0}, {"value", 162}},
		{{"year", 2017}, {"offenseclass", 1}, {"value", 1277}},
		{{"year", 2017}, {"offenseclass", 2}, {"value", 81}},
		{{"year", 2018}, {"offenseclass", 0}, {"value", 421}},
		{{"year", 2018}, {"offenseclass", 1}, {"value", 1253}},
		{{"year", 2018}, {"offenseclass", 2}, {"value", 121}},
	}

	mr := make(map[string]data.RefTable)
	mr["RefOffenseClass"] = data.RefTable{
		{0, "felony"},
		{1, "misdemeanor"},
		{2, "unknown"},
	}

	s := new(Storage)
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
