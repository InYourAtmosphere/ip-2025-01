package main

import (
	f "fmt"
)

func main() {

	var cima, baixo int
	var h float64

	for i := 0; i < 50; i++ {

		cima = 1 + 2*i

		baixo = 1 + i

		h = float64(cima/baixo) + h

	}

	f.Printf("%f\n", h)
}
