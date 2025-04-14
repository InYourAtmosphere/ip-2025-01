package main

import (
	f "fmt"
)

func main() {

	var horas, totalPagar int

	f.Scanf("%d", &horas)

	totalPagar = 10*(horas/3) + 5*(horas%3)

	f.Printf("O VALOR A PAGAR E = %.2d\n", totalPagar)

}
