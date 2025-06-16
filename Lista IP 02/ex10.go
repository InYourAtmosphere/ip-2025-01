package main

import (
	"fmt"
)

func main() {
	var destino, pernaViagem int
	var valor float64

	// Laço para entrada do destino
	for {
		fmt.Println("Qual o destino?")
		fmt.Println("1 - Norte")
		fmt.Println("2 - Nordeste")
		fmt.Println("3 - Centro-Oeste")
		fmt.Println("4 - Região Sul")
		fmt.Print("Escolha: ")
		fmt.Scanf("%d", &destino)

		if destino >= 1 && destino <= 4 {
			break
		}
		fmt.Println("--------- Número incorreto ----------")
	}

	// Laço para entrada da perna da viagem

	fmt.Println("Perna viagem")
	fmt.Println("1 - Ida")
	fmt.Println("2 - Ida e Volta")
	fmt.Print("Escolha: ")

	for {
		fmt.Scanf("%d", &pernaViagem)
		if pernaViagem == 1 || pernaViagem == 2 {
			break
		}
		fmt.Println("--------- Número incorreto ----------")
	}

	// Definir valor com base no destino e tipo de viagem
	switch destino {
	case 1: // Norte
		if pernaViagem == 1 {
			valor = 500
		} else {
			valor = 900
		}
	case 2: // Nordeste
		if pernaViagem == 1 {
			valor = 350
		} else {
			valor = 650
		}
	case 3: // Centro-Oeste
		if pernaViagem == 1 {
			valor = 350
		} else {
			valor = 600
		}
	case 4: // Região Sul
		if pernaViagem == 1 {
			valor = 300
		} else {
			valor = 600
		}
	}

	fmt.Printf("\nDestino: %d\nPerna de viagem: %d\nValor: R$ %.2f\n", destino, pernaViagem, valor)
}
