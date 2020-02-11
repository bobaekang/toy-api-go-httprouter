package data

// Repository provides access to arrests data storage.
type Repository interface {
	GetTable(string) Table
	GetRefTable(string) RefTable
}

// Service provides operations for arrests data.
type Service interface {
	GetTable(string) Table
	GetRefTable(string) RefTable
}

type service struct {
	r Repository
}

// NewService creates a arrests service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// GetTable returns a data for the given table name.
func (s *service) GetTable(tableName string) Table {
	return s.r.GetTable(tableName)
}

// GetRefTable returns a reference for the given table name.
func (s *service) GetRefTable(tableName string) RefTable {
	return s.r.GetRefTable(tableName)
}
