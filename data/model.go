package data

import (
	"bytes"
	"encoding/json"
	"sort"
)

// Variable models a variable name-value pair
type Variable struct {
	Name  string
	Value int
}

// Row models a pairing of Variables and a value
type Row []Variable

// Table models a collection of Rows
type Table []Row

// Filter returns a new Table filtered with a specified condition
func (table Table) Filter(by string, matchIf string, value int) Table {
	for i := 0; i < len(table); i++ {
		match := false

		for _, v := range table[i] {
			switch matchIf {
			case "==":
				if v.Name == by && v.Value == value {
					match = true
				}
			case "<=":
				if v.Name == by && v.Value <= value {
					match = true
				}
			case ">=":
				if v.Name == by && v.Value >= value {
					match = true
				}
			case "<":
				if v.Name == by && v.Value < value {
					match = true
				}
			case ">":
				if v.Name == by && v.Value > value {
					match = true
				}
			}
		}

		if !match {
			table = append(table[:i], table[i+1:]...)
			i--
		}
	}

	return table
}

// Select implements selecting Variables by name operation for Table type
func (table Table) Select(varNames ...string) Table {
	selected := make(Table, len(table))

	for i := range table {
		var row Row

		for _, v := range table[i] {
			for _, varName := range varNames {
				if v.Name == varName {
					row = append(row, v)
				}
			}
		}

		selected[i] = row
	}

	return selected
}

// SortBy returns a new sorted Table
func (table Table) SortBy(by string, order string) Table {
	sorted := table

	sort.SliceStable(sorted, func(i, j int) bool {
		var iVal, jVal int

		for _, v := range sorted[i] {
			if v.Name == by {
				iVal = v.Value
			}
		}

		for _, v := range sorted[j] {
			if v.Name == by {
				jVal = v.Value
			}
		}

		var less bool

		switch order {
		case "asc":
			less = iVal < jVal
		case "desc":
			less = iVal > jVal
		}

		return less
	})

	return sorted
}

// MarshalJSON implements custom JSON marshaler for Table
func (table Table) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("[")

	for i, row := range table {
		if i != 0 {
			buf.WriteString(",")
		}

		buf.WriteString("{")

		// marshal Variables
		for j, variable := range row {
			if j != 0 {
				buf.WriteString(",")
			}

			err := writeJSONProp(&buf, variable.Name, variable.Value)
			if err != nil {
				return nil, err
			}
		}

		buf.WriteString("}")
	}

	buf.WriteString("]")

	return buf.Bytes(), nil
}

func writeJSONProp(buf *bytes.Buffer, key string, val int) error {
	k, err := json.Marshal(key)
	if err != nil {
		return err
	}

	v, err := json.Marshal(val)
	if err != nil {
		return err
	}

	buf.Write(k)
	buf.WriteString(":")
	buf.Write(v)

	return nil
}
