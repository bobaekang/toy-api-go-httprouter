package main

type ArrestsAllOne struct {
	year  int
	value int
}

type ArrestsByOffenseClassOne struct {
	year         int
	offenseclass int
	value        int
}

type ArrestsAll []ArrestsAllOne

type ArrestsByOffenseClass []ArrestsByOffenseClassOne
