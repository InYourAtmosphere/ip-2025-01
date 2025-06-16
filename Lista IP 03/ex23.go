package main

import (
	f "fmt"
	m "math"
)

func main() {

	var n int
	var soma float64
	f.Scanln(&n)

	for i := 0; i < n; i++ {

		soma = m.Pow(-1, float64(i))*float64(1000-i*3)/1 + float64(i)
	}

	f.Print(" %.2f ", soma)
}
