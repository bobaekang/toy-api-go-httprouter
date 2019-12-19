package arrests

// All models a row in Arrests table.
type All struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// ByOffenseClass models a row in ArrestsByOffenseClass table.
type ByOffenseClass struct {
	Year         int `json:"year"`
	OffenseClass int `json:"offenseclass"`
	Value        int `json:"value"`
}
