package main

type ArrestsAllOne struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type ArrestsByOffenseClassOne struct {
	Year         int `json:"year"`
	OffenseClass int `json:"offenseclass"`
	Value        int `json:"value"`
}

type ArrestsAll []ArrestsAllOne

type ArrestsByOffenseClass []ArrestsByOffenseClassOne
