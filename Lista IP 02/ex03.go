package main

import (
	f "fmt"
)

func main() {

	var numero1, numero2 float64

	f.Scanf("%f %f", &numero1, &numero2)

	if numero1+numero2 > 20 {
		resultado := numero1 + numero2 + 8
		f.Printf("%.2f\n", resultado)
	} else {

		resultado := numero1 + numero2

		f.Printf("%.2f\n", resultado)
	}

}
