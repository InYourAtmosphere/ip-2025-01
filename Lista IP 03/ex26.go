package main

import (
	f "fmt"
)

func main() {

	var produto, soma, cima float64

	produto = 1
	soma = 0.0
	for i := 0; i < 20; i++ {

		for q := 1; q <= i; q++ {

			produto = produto * float64(q)
		}
		cima = 100 - float64(i)
		parcial := cima / produto
		f.Printf(" %4.16f ", parcial)
		soma = soma + cima/produto

	}
	f.Print("\n soma final %f\n", soma)
}
