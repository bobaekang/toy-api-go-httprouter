package model

type Repository interface {
	GetArrestsAll() []ArrestsAll
	GetArrestsByOffenseClass() []ArrestsByOffenseClass
}

type Service interface {
	GetArrestsAll() []ArrestsAll
	GetArrestsByOffenseClass() []ArrestsByOffenseClass
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetArrestsAll() []ArrestsAll {
	return s.r.GetArrestsAll()
}

func (s *service) GetArrestsByOffenseClass() []ArrestsByOffenseClass {
	return s.r.GetArrestsByOffenseClass()
}
