package main

import (
	f "fmt"
	m "math"
)

func main() {

	var A, B, C float64

	f.Scanf("%f %f %f", &A, &B, &C)

	delta := B*B - 4*A*C

	if A == 0 {
		return
	}
	if delta > 0 {

		raiz := -(B - m.Pow(delta, 0.5)) / (2 * A)
		raizDois := -(B + m.Pow(delta, 0.5)) / (2 * A)
		f.Printf("raiz 1= %f, raiz 2 = %f", raiz, raizDois)
	} else if delta == 0 {

		raizUnica := -(B) / (2 * A)
		f.Printf("raiz unica %", raizUnica)
	} else if delta < 0 {
		f.Printf("raiz imaginaria")
	}
}
