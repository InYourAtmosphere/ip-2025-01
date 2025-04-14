package main

import (
	"fmt"
)

func main() {

	var salario float64

	fmt.Scanf("%f", &salario)

	if salario <= 300 {
		reajuste := salario * 1.5

		fmt.Printf("SALARIO COM REAJUSTE = %.2f", reajuste)
	} else {

		reajuste := salario * 1.3

		fmt.Printf("SALARIO COM REAJUSTE = %.2f", reajuste)
	}
}
