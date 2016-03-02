package main

import (
	"math"
)

func Extend(slice []float64, element float64) []float64 {
	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]float64, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func Append(slice []float64, items ...float64) []float64 {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

func ExpArray(arr []float64) []float64 {
	// Golang implementation of Python numpy.exp
	var expArr []float64
	for i := range arr {
		expArr = Append(expArr, math.Exp(arr[i]))
	}
	return expArr
}

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// func sigmoid_out2deriv(out) {
//     return out * (1 - out)
// }

// func tanh(x) {
//     return np.tanh(x)
// }

// func tanh_out2deriv(out) {
//     return (1 - out**2)
// }

// func relu(x, deriv=False) {
//     if(deriv):
//         return int(x >= 0)
//     return math.Max(0, x)
// }
