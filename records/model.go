package records

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

// RecordMap models Record in the flattened map format
type RecordMap map[string]int

// Records models a collection of Records
type Records []Record

// Flatten wrangles Records into a slice of flat RecordMaps
func (aa Records) Flatten() []RecordMap {
	var recordMaps []RecordMap

	for _, a := range aa {
		recordMap := make(RecordMap)

		for _, g := range a.Groups {
			recordMap[g.Name] = g.Value
		}
		recordMap["value"] = a.Value

		recordMaps = append(recordMaps, recordMap)
	}

	return recordMaps
}
