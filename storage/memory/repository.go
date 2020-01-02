package memory

import (
	"github.com/bobaekang/toy-api-go-httprouter/data"
)

type Storage struct {
	ArrestsAll            data.Table
	ArrestsByOffenseClass data.Table
}

func NewStorage() *Storage {
	s := new(Storage)
	s.ArrestsAll = data.Table{
		{
			Variables: []data.Variable{{"year", 2017}},
			Value:  1820,
		},
		{
			Variables: []data.Variable{{"year", 2018}},
			Value:  1795,
		},
	}
	s.ArrestsByOffenseClass = data.Table{
		{
			Variables: []data.Variable{{"year", 2017}, {"offenseclass", 0}},
			Value:  162,
		},
		{
			Variables: []data.Variable{{"year", 2017}, {"offenseclass", 1}},
			Value:  1277,
		},
		{
			Variables: []data.Variable{{"year", 2017}, {"offenseclass", 2}},
			Value:  81,
		},
		{
			Variables: []data.Variable{{"year", 2018}, {"offenseclass", 0}},
			Value:  421,
		},
		{
			Variables: []data.Variable{{"year", 2018}, {"offenseclass", 1}},
			Value:  1253,
		},
		{
			Variables: []data.Variable{{"year", 2018}, {"offenseclass", 2}},
			Value:  121,
		},
	}

	return s
}

func (s *Storage) GetArrestsAll() data.Table {
	return s.ArrestsAll
}

func (s *Storage) GetArrestsByOffenseClass() data.Table {
	return s.ArrestsByOffenseClass
}
