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

func SigmoidOut2Deriv(out float64) float64 {
	return out * (1 - out)
}

func Tanh(arr []float64) []float64 {
	var tanhArr []float64
	for i := range arr {
		tanhArr = Append(tanhArr, math.Tanh(arr[i]))
	}
	return tanhArr
}

func TanhOut2Deriv(out float64) float64 {
	return (1.0 - math.Pow(out, 2))
}

// func relu(x, deriv=False) {
//     if(deriv):
//         return int(x >= 0)
//     return math.Max(0, x)
// }
