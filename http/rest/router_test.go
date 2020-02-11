package rest

import "testing"

func TestToPath(t *testing.T) {
	got := toPath("Arrests")
	expected := "/arrests"
	if got != expected {
		t.Error(
			"toPath:",
			"\n     got ", got,
			"\nexpected ", expected,
		)
	}

	got = toPath("ArrestsByOffenseClass")
	expected = "/arrests/by-offense-class"
	if got != expected {
		t.Error(
			"toPath:",
			"\n     got ", got,
			"\nexpected ", expected,
		)
	}

	got = toPath("RefOffenseClass")
	expected = "/ref/offense-class"
	if got != expected {
		t.Error(
			"toPath:",
			"\n     got ", got,
			"\nexpected ", expected,
		)
	}
}
