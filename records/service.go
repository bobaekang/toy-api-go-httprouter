package records

// Repository provides access to arrests data storage.
type Repository interface {
	GetArrestsAll() Records
	GetArrestsByOffenseClass() Records
}

// Service provides operations for arrests data.
type Service interface {
	GetArrestsAll() Records
	GetArrestsByOffenseClass() Records
}

type service struct {
	r Repository
}

// NewService creates a arrests service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// GetArrestsAll returns arrests data.
func (s *service) GetArrestsAll() Records {
	return s.r.GetArrestsAll()
}

// GetArrestsAll returns arrests by offense class data.
func (s *service) GetArrestsByOffenseClass() Records {
	return s.r.GetArrestsByOffenseClass()
}
