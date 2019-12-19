package arrests

type All struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type ByOffenseClass struct {
	Year         int `json:"year"`
	OffenseClass int `json:"offenseclass"`
	Value        int `json:"value"`
}
