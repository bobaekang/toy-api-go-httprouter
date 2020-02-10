package data

// Repository provides access to arrests data storage.
type Repository interface {
	GetTable(string) Table
}

// Service provides operations for arrests data.
type Service interface {
	GetArrestsAll() Table
	GetArrestsByOffenseClass() Table
}

type service struct {
	r Repository
}

// NewService creates a arrests service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// GetArrestsAll returns arrests data.
func (s *service) GetArrestsAll() Table {
	return s.r.GetTable("ArrestsAll")
}

// GetArrestsAll returns arrests by offense class data.
func (s *service) GetArrestsByOffenseClass() Table {
	return s.r.GetTable("ArrestsByOffenseClass")
}
