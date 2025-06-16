package main

import "fmt"

func main() {
	var numeros []int

	fmt.Println("Digite os números inteiros que deseja verificar (digite 0 ou um número negativo para finalizar a entrada):")

	for {
		var numero int
		fmt.Print("Digite um número: ")
		fmt.Scan(&numero)

		if numero <= 0 {
			break
		}

		numeros = append(numeros, numero)
	}

	fmt.Println("\nResultados:")
	for _, num := range numeros {
		if num < 0 {
			fmt.Printf("%d não é um quadrado perfeito.\n", num)
		} else if num == 0 || num == 1 {
			fmt.Printf("%d é um quadrado perfeito.\n", num)
		} else {
			encontrouRaiz := false
			for i := 1; i*i <= num; i++ {
				if i*i == num {
					fmt.Printf("%d é um quadrado perfeito.\n", num)
					encontrouRaiz = true
					break
				}
			}
			if !encontrouRaiz {
				fmt.Printf("%d não é um quadrado perfeito.\n", num)
			}
		}
	}

	fmt.Println("Programa encerrado.")
}
