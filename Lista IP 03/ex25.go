package main

import (
	f "fmt"
	m "math"
)

func main() {

	var soma float64

	for i := 0.0; i < 25; i++ {

		soma = soma + m.Pow(-1, i)*(m.Pow(2, i)/m.Pow(25-i, 2))
	}

	f.Printf(" %3.3f ", soma)
}
