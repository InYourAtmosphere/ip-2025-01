package main

import (
	"fmt"
)

func main() {

	var popular, geral, arquibancada, cadeiras, publico, somaReceita, a, b, c, d float64

	fmt.Scanf("%f %f %f %f %f", publico, popular, geral, arquibancada, cadeiras)

	a = ((publico * popular) / 100) * 1
	b = ((publico * geral) / 100) * 5
	c = ((publico * arquibancada) / 100) * 10
	d = ((publico * cadeiras) / 100) * 20

	somaReceita = a + b + c + d //poderia escrever tudo em uma fórmula mas ficaria muito carregado

	fmt.Printf("A RENDA DO JOGO %f", somaReceita)
}
