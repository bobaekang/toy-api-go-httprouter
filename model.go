package main

type ArrestsAll struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type ArrestsByOffenseClass struct {
	Year         int `json:"year"`
	OffenseClass int `json:"offenseclass"`
	Value        int `json:"value"`
}
