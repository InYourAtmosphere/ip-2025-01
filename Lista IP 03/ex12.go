package main

import (
	f "fmt"
	m "math"
)

func main() {

	var soma, x, produto float64
	soma = 0.0
	produto = 1.0
	f.Scanln(&x)
	for i := 1.0; i <= 20; i++ {
		produto = produto * i
		soma = soma + m.Pow(-1, i)/produto
	}

	resultadoFinal := -1 * (x * soma)

	f.Printf("%.2f\n", resultadoFinal)
}
