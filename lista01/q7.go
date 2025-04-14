package main

import (
	"fmt"
)

func main() {

	var FAHRENHEIT, POLEGADAS float64
	fmt.Scanf("%f", &FAHRENHEIT)
	fmt.Scanf("%f", &POLEGADAS)

	converteTemperatura := 5 * (FAHRENHEIT - 32) / 9

	converteChuva := POLEGADAS * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA E = %.2f\n", converteChuva, converteTemperatura)
}
