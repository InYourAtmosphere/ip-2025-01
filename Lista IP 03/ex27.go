package main

import (
	f "fmt"
	m "math"
)

func main() {

	var x, soma, cima, baixo float64

	f.Scanln(&x)

	baixo = 2

	for i := 1.0; i <= 20.0; i++ {

		for q := 1.0; q <= i; q++ {

			baixo = baixo * q
		}

		soma = soma + (m.Pow(-1, i)*m.Pow(x, 2*i))/baixo + 1.0

	}

	f.Printf("%f\n", soma)

}
