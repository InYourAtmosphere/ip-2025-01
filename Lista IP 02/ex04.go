package main

import (
	f "fmt"
	m "math"
)

func main() {

	var numero float64

	f.Scanf("%f", &numero)

	if numero >= 0 {

		resultado := m.Pow(numero, 0.5)

		f.Printf("%.2f", resultado)
	} else {

		resultado := m.Pow(numero, 2)
		f.Printf("%.2f", resultado)
	}
}
