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
		t.Error("actualOut from ExpArray did not equal the expectedOut")
	}
}

func TestSigmoid(t *testing.T) {
	testIn := 5.0
	expectedOut := 0.99330714907571527
	actualOut := Sigmoid(testIn)
	if actualOut != expectedOut {
		t.Error("actualOut from Sigmoid did not equal the expectedOut")
	}
}

func TestSigmoidOut2Deriv(t *testing.T) {
	testIn := 5.0
	expectedOut := -20.0
	actualOut := SigmoidOut2Deriv(testIn)
	if actualOut != expectedOut {
		t.Error("actualOut from SigmoidOut2Deriv did not equal the expectedOut")
	}
}

func TestTanh(t *testing.T) {
	testIn := []float64{0.0, 1.2, 3.3}
	expectedOut := []float64{0., 0.83365461, 0.99728296}
	actualOut := Tanh(testIn)
	if testEq(actualOut, expectedOut) {
		t.Error("actualOut from Tanh did not equal the expectedOut")
	}
}
