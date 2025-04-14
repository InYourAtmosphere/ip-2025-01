package main

import (
	f "fmt"
)

func main() {

	var A, B, C float64

	f.Scanf("%f", &A)
	f.Scanf("%f", &B)
	f.Scanf("%f", &C)

	Delta := B*B - 4*A*C

	f.Print("O VALOR DE DELTA E = %.2f\n", Delta)

}
