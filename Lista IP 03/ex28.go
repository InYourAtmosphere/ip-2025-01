package main

import (
	f "fmt"
	m "math"
)

func main() {

	var soma, p float64

	for i := 0.0; i < 51; i++ {

		soma = soma + m.Pow(-1, i)*1/m.Pow(2*i+1, 3)
	}

	p = m.Cbrt(soma * 32)

	f.Printf("%.9f\n", p)

}
