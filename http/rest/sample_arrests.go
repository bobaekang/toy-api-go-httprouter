package rest

import "github.com/bobaekang/toy-api-go-httprouter/model"

var SampleDataArrestsAll = []model.ArrestsAll{
	{
		Year:  2017,
		Value: 1820,
	},
	{
		Year:  2018,
		Value: 1795,
	},
}

var SampleDataArrestsByOffenseClass = []model.ArrestsByOffenseClass{
	{
		Year:         2017,
		OffenseClass: 0,
		Value:        162,
	},
	{
		Year:         2017,
		OffenseClass: 1,
		Value:        1277,
	},
	{
		Year:         2017,
		OffenseClass: 2,
		Value:        81,
	},
	{
		Year:         2018,
		OffenseClass: 0,
		Value:        421,
	},
	{
		Year:         2018,
		OffenseClass: 1,
		Value:        1253,
	},
	{
		Year:         2018,
		OffenseClass: 2,
		Value:        121,
	},
}
