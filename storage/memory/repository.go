package memory

import (
	"github.com/bobaekang/toy-api-go-httprouter/records"
)

type Storage struct {
	ArrestsAll            records.Records
	ArrestsByOffenseClass records.Records
}

func NewStorage() *Storage {
	s := new(Storage)
	s.ArrestsAll = records.Records{
		{
			Groups: []records.Group{{"year", 2017}},
			Value:  1820,
		},
		{
			Groups: []records.Group{{"year", 2018}},
			Value:  1795,
		},
	}
	s.ArrestsByOffenseClass = records.Records{
		{
			Groups: []records.Group{{"year", 2017}, {"offenseclass", 0}},
			Value:  162,
		},
		{
			Groups: []records.Group{{"year", 2017}, {"offenseclass", 1}},
			Value:  1277,
		},
		{
			Groups: []records.Group{{"year", 2017}, {"offenseclass", 2}},
			Value:  81,
		},
		{
			Groups: []records.Group{{"year", 2018}, {"offenseclass", 0}},
			Value:  421,
		},
		{
			Groups: []records.Group{{"year", 2018}, {"offenseclass", 1}},
			Value:  1253,
		},
		{
			Groups: []records.Group{{"year", 2018}, {"offenseclass", 2}},
			Value:  121,
		},
	}

	return s
}

func (s *Storage) GetArrestsAll() records.Records {
	return s.ArrestsAll
}

func (s *Storage) GetArrestsByOffenseClass() records.Records {
	return s.ArrestsByOffenseClass
}
