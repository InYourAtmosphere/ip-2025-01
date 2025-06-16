package main

import (
	"fmt"
)

func main() {

	var numero, soma, fundo, topo, mediaPares, percentil, media float64
	var contador, impar, q int

	fundo = 1e18

	for i := 0; i > -10; i++ {
		fmt.Print("Digite um número (30000 para sair): ")
		fmt.Scanln(&numero)

		if numero == 30000 {
			break
		}

		soma = soma + numero

		if i == 0 || numero > topo {
			topo = numero
		}

		if i == 0 || numero < fundo {
			fundo = numero
		}

		if int(numero)%2 == 0 {
			mediaPares += numero
			q++
		}

		contador++
		impar = contador - q // total menos pares

		if contador > 0 {
			media = soma / float64(contador)
			percentil = float64(impar) / float64(contador)
		}
	}

	var mediaParesFinal float64
	if q > 0 {
		mediaParesFinal = mediaPares / float64(q)
	}

	fmt.Printf("Soma dos números: %.2f\n", soma)
	fmt.Printf("Quantidade de números digitados: %d\n", contador)
	fmt.Printf("Média dos números: %.2f\n", media)
	fmt.Printf("Maior número digitado: %.2f\n", topo)
	fmt.Printf("Menor número digitado: %.2f\n", fundo)
	fmt.Printf("Média dos números pares: %.2f\n", mediaParesFinal)
	fmt.Printf("Percentagem dos números ímpares entre todos os números digitados: %.2f%%\n", percentil*100)
}
