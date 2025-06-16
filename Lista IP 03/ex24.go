package main

import (
	f "fmt"
	m "math"
)

func main() {

	var baixo, soma float64

	for a := 0.0; a < 6.3; a = a + 0.1 {
		soma = a
		for i := 1; i < 4; i++ {

			baixo = 1.0

			for q := 1; q <= 2*int(i)+1; q++ {
				baixo = baixo * float64(q)
			}
			soma = (m.Pow(-1, float64(i))*m.Pow(a, 2*float64(i)+1))/baixo + soma

		}

		f.Printf(" %.3f ", soma)

	}
}
