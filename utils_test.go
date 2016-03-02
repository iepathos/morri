package main

import "testing"

func testEq(a, b []float64) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestExpArray(t *testing.T) {
	testIn := []float64{0.0, 1.2, 3.3}
	expectedOut := []float64{1., 3.32011692, 27.11263892}
	actualOut := ExpArray(testIn)
	if testEq(actualOut, expectedOut) {
		t.Error("actualOut from ExpArray not equal to expectedOut")
	}
}
