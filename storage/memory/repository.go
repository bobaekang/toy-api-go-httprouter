package memory

import (
	"github.com/bobaekang/toy-api-go-httprouter/data"
)

type Storage struct {
	tables map[string]data.Table
}

func NewStorage() *Storage {
	m := make(map[string]data.Table)
	m["Arrests"] = data.Table{
		{{"year", 2017}, {"value", 1820}},
		{{"year", 2018}, {"value", 1795}},
	}
	m["ArrestsByOffenseClass"] = data.Table{
		{{"year", 2017}, {"offenseclass", 0}, {"value", 162}},
		{{"year", 2017}, {"offenseclass", 1}, {"value", 1277}},
		{{"year", 2017}, {"offenseclass", 2}, {"value", 81}},
		{{"year", 2018}, {"offenseclass", 0}, {"value", 421}},
		{{"year", 2018}, {"offenseclass", 1}, {"value", 1253}},
		{{"year", 2018}, {"offenseclass", 2}, {"value", 121}},
	}

	s := new(Storage)
	s.tables = m

	return s
}

func (s *Storage) GetTable(tableName string) data.Table {
	return s.tables[tableName]
}
