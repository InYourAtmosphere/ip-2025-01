package main

import (
	f "fmt"
)

func main() {

	var A, B, C, MAIOR, INTER, MENOR float64

	f.Scanf("%f %f %f", &A, &B, &C)

	if A < B && B < C {

		A = MAIOR
		B = INTER
		C = MENOR

		f.Printf("%f %f %f", MAIOR, INTER, MENOR)
	}
}
