package main

import (
	f "fmt"
	m "math"
)

func main() {

	var salarioCarlos, somaJoao, somaCarlos, mes float64

	f.Scanf("%f", &salarioCarlos)

	//obs: a formula independe dos salários iniciais de calos ou joão

	somaJoao = 0
	somaCarlos = 0

	for i := 0.0; somaJoao <= somaCarlos; i++ {

		somaJoao = salarioCarlos * (m.Pow(1.05, i) - 1) / 0.15

		somaCarlos = salarioCarlos * (m.Pow(1.02, i) - 1) / 0.02

		f.Printf("soma Joao: %.2f	soma Carlos: %.2f mes: %.0f\n", somaJoao, somaCarlos, i)
		mes = i
	}

	f.Printf("mes %.2f\n", mes+2)
}
