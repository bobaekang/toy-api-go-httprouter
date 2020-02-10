package data

import (
	"fmt"
	"reflect"
	"testing"
)

func getSampleData() Table {
	return Table{
		{{"colA", 1}, {"colB", 1}, {"value", 1}},
		{{"colA", 1}, {"colB", 2}, {"value", 2}},
		{{"colA", 2}, {"colB", 1}, {"value", 3}},
		{{"colA", 2}, {"colB", 2}, {"value", 4}},
		{{"colA", 3}, {"colB", 1}, {"value", 5}},
		{{"colA", 3}, {"colB", 2}, {"value", 6}},
	}
}

func TestFilter(t *testing.T) {
	got1 := getSampleData()
	got1 = got1.Filter("colA", "==", 3)
	expected1 := Table{
		{{"colA", 3}, {"colB", 1}, {"value", 5}},
		{{"colA", 3}, {"colB", 2}, {"value", 6}},
	}

	if !reflect.DeepEqual(got1, expected1) && len(got1) == len(expected1) {
		t.Error(
			"Filter: colA == 3",
			"\n     got ", got1,
			"\nexpected ", expected1,
		)
	}

	got2 := getSampleData()
	got2 = got2.Filter("colB", "<", 2)
	expected2 := Table{
		{{"colA", 1}, {"colB", 1}, {"value", 1}},
		{{"colA", 2}, {"colB", 1}, {"value", 3}},
		{{"colA", 3}, {"colB", 1}, {"value", 5}},
	}

	if !reflect.DeepEqual(got2, expected2) {
		t.Error(
			"Filter: colB < 2",
			"\n     got ", got2,
			"\nexpected ", expected2,
		)
	}

	got3 := getSampleData()
	got3 = got3.Filter("colA", ">=", 2)
	expected3 := Table{
		{{"colA", 2}, {"colB", 1}, {"value", 3}},
		{{"colA", 2}, {"colB", 2}, {"value", 4}},
		{{"colA", 3}, {"colB", 1}, {"value", 5}},
		{{"colA", 3}, {"colB", 2}, {"value", 6}},
	}

	if !reflect.DeepEqual(got3, expected3) {
		t.Error(
			"Filter: colA >= 2",
			"\n     got ", got3,
			"\nexpected ", expected3,
		)
	}
}

func TestSelect(t *testing.T) {
	got := getSampleData()
	got = got.Select("colA", "value")
	expected := Table{
		{{"colA", 1}, {"value", 1}},
		{{"colA", 1}, {"value", 2}},
		{{"colA", 2}, {"value", 3}},
		{{"colA", 2}, {"value", 4}},
		{{"colA", 3}, {"value", 5}},
		{{"colA", 3}, {"value", 6}},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"Select: colA, value",
			"\n     got ", got,
			"\nexpected ", expected,
		)
	}
}

func TestSortBy(t *testing.T) {
	got := getSampleData()
	got = got.SortBy("colA", "asc")
	got = got.SortBy("colB", "desc")
	expected := Table{
		{{"colA", 1}, {"colB", 2}, {"value", 2}},
		{{"colA", 2}, {"colB", 2}, {"value", 4}},
		{{"colA", 3}, {"colB", 2}, {"value", 6}},
		{{"colA", 1}, {"colB", 1}, {"value", 1}},
		{{"colA", 2}, {"colB", 1}, {"value", 3}},
		{{"colA", 3}, {"colB", 1}, {"value", 5}},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"SortBy: colA then by DESC(colB)",
			"\n     got ", got,
			"\nexpected ", expected,
		)
	}
}

func TestJSONMarshaler(t *testing.T) {
	data := getSampleData()
	j, err := data.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	expected := `[{"colA":1,"colB":1,"value":1},{"colA":1,"colB":2,"value":2},{"colA":2,"colB":1,"value":3},{"colA":2,"colB":2,"value":4},{"colA":3,"colB":1,"value":5},{"colA":3,"colB":2,"value":6}]`

	if got := string(j); expected != got {
		t.Error(
			"MarshalJSON",
			"\n     got ", got,
			"\nexpected ", expected,
		)
	}
}
