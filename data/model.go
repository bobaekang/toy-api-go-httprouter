package data

import (
	"bytes"
	"encoding/json"
)

// Variable models a variable name-value pair
type Variable struct {
	Name  string
	Value int
}

// Row models a pairing of Variables and a value
type Row struct {
	Variables []Variable
	Value  int
}

// Table models a collection of Rows
type Table []Row

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
		for j, variable := range row.Variables {
			if j != 0 {
				buf.WriteString(",")
			}

			err := writeJSONProp(&buf, variable.Name, variable.Value)
			if err != nil {
				return nil, err
			}
		}

		buf.WriteString(",")

		// marshal Value
		err := writeJSONProp(&buf, "value", row.Value)
		if err != nil {
			return nil, err
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
