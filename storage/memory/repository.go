package memory

import "github.com/bobaekang/toy-api-go-httprouter/model"

type Storage struct {
	ArrestsAll            []model.ArrestsAll
	ArrestsByOffenseClass []model.ArrestsByOffenseClass
}

func NewStorage() *Storage {
	s := new(Storage)
	s.ArrestsAll = []model.ArrestsAll{
		{
			Year:  2017,
			Value: 1820,
		},
		{
			Year:  2018,
			Value: 1795,
		},
	}
	s.ArrestsByOffenseClass = []model.ArrestsByOffenseClass{
		{
			Year:         2017,
			OffenseClass: 0,
			Value:        162,
		},
		{
			Year:         2017,
			OffenseClass: 1,
			Value:        1277,
		},
		{
			Year:         2017,
			OffenseClass: 2,
			Value:        81,
		},
		{
			Year:         2018,
			OffenseClass: 0,
			Value:        421,
		},
		{
			Year:         2018,
			OffenseClass: 1,
			Value:        1253,
		},
		{
			Year:         2018,
			OffenseClass: 2,
			Value:        121,
		},
	}

	return s
}

func (s *Storage) GetArrestsAll() []model.ArrestsAll {
	return s.ArrestsAll
}

func (s *Storage) GetArrestsByOffenseClass() []model.ArrestsByOffenseClass {
	return s.ArrestsByOffenseClass
}
