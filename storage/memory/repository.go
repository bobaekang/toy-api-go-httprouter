package memory

import "github.com/bobaekang/toy-api-go-httprouter/arrests"

type Storage struct {
	ArrestsAll            []arrests.All
	ArrestsByOffenseClass []arrests.ByOffenseClass
}

func NewStorage() *Storage {
	s := new(Storage)
	s.ArrestsAll = []arrests.All{
		{
			Year:  2017,
			Value: 1820,
		},
		{
			Year:  2018,
			Value: 1795,
		},
	}
	s.ArrestsByOffenseClass = []arrests.ByOffenseClass{
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

func (s *Storage) GetArrestsAll() []arrests.All {
	return s.ArrestsAll
}

func (s *Storage) GetArrestsByOffenseClass() []arrests.ByOffenseClass {
	return s.ArrestsByOffenseClass
}
