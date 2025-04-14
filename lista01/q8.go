package main

import (
	"fmt"
)

const Pi = 3.14159

func main() {

	var raio, altura float64

	fmt.Scanf("%f", raio)
	fmt.Scanf("%f", altura)

	// custo do m² 100

	areaBase := (Pi * raio * raio)

	areaLateral := 2 * Pi * raio * altura

	custo := 2 * (2*areaBase + areaLateral)

	fmt.Print("O VALOR DO CUSTO E = %.2f\n", custo)

}
