package arrests

type Repository interface {
	GetArrestsAll() []All
	GetArrestsByOffenseClass() []ByOffenseClass
}

type Service interface {
	GetArrestsAll() []All
	GetArrestsByOffenseClass() []ByOffenseClass
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetArrestsAll() []All {
	return s.r.GetArrestsAll()
}

func (s *service) GetArrestsByOffenseClass() []ByOffenseClass {
	return s.r.GetArrestsByOffenseClass()
}
