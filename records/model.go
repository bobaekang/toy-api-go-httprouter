package records

import (
	"bytes"
	"encoding/json"
)

// Group models a group name-value pair
type Group struct {
	Name  string
	Value int
}

// Record models a pairing of Groups and a value
type Record struct {
	Groups []Group
	Value  int
}

// Records models a collection of Records
type Records []Record

// MarshalJSON implements custom JSON marshaler for Records
func (aa Records) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("[")

	for i, a := range aa {
		if i != 0 {
			buf.WriteString(",")
		}

		buf.WriteString("{")

		// marshal Groups
		for j, g := range a.Groups {
			if j != 0 {
				buf.WriteString(",")
			}

			err := writeJSONProp(&buf, g.Name, g.Value)
			if err != nil {
				return nil, err
			}
		}

		buf.WriteString(",")

		// marshal Value
		err := writeJSONProp(&buf, "value", a.Value)
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
