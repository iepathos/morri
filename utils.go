package main

import (
	"math"
)

// Extends a float64 by a float64 array
func Extend(slice []float64, element float64) []float64 {
	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]float64, n, 2*n+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// Appends a float64 to a float64 array
func Append(slice []float64, items ...float64) []float64 {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

// Sum returns the sum of the elements
// in a float64 array. Golang implementation of Python numpy.sum
func Sum(arr []float64) float64 {
	sum := 0.0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

// ExpArray returns the exponential array
// of the given []float64 array.
// Golang implementation of Python numpy.exp
func ExpArray(arr []float64) []float64 {
	var expArr []float64
	for i := range arr {
		expArr = Append(expArr, math.Exp(arr[i]))
	}
	return expArr
}

// Sigmoid
func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// SigmoidOut2Deriv
func SigmoidOut2Deriv(out float64) float64 {
	return out * (1 - out)
}

// Tanh is a Golang implementation of Python numpy.tanh
func Tanh(arr []float64) []float64 {
	var tanhArr []float64
	for i := range arr {
		tanhArr = Append(tanhArr, math.Tanh(arr[i]))
	}
	return tanhArr
}

// TanhOut2Deriv
func TanhOut2Deriv(out float64) float64 {
	return (1.0 - math.Pow(out, 2))
}

// Relu
func Relu(x float64, deriv bool) float64 {
	if deriv {
		if x >= 0 {
			return 1.0
		}
		return 0.0
	}
	return math.Max(0, x)
}
