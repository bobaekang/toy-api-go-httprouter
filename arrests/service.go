package arrests

// Repository provides access to arrests data storage.
type Repository interface {
	GetArrestsAll() []All
	GetArrestsByOffenseClass() []ByOffenseClass
}

// Service provides operations for arrests data.
type Service interface {
	GetArrestsAll() []All
	GetArrestsByOffenseClass() []ByOffenseClass
}

type service struct {
	r Repository
}

// NewService creates a arrests service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// GetArrestsAll returns arrests data.
func (s *service) GetArrestsAll() []All {
	return s.r.GetArrestsAll()
}

// GetArrestsAll returns arrests by offense class data.
func (s *service) GetArrestsByOffenseClass() []ByOffenseClass {
	return s.r.GetArrestsByOffenseClass()
}
