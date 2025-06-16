package main

import (
	"fmt"
	"strings"
)

func main() {
	var dia int
	var precoNormal, precoFinal float64
	var categoria string

	fmt.Print("Informe o dia da semana (1-Dom, 2-Seg, ..., 7-Sáb): ")
	fmt.Scanf("%d", &dia)

	fmt.Print("Informe o preço normal do DVD (R$): ")
	fmt.Scanf("%f", &precoNormal)

	fmt.Print("Informe a categoria (comum ou lançamento): ")
	fmt.Scanf("%s", &categoria)

	categoria = strings.ToLower(categoria) // normaliza para evitar erros com maiúsculas

	// Aplica desconto de 40% para segunda, terça e quinta
	if dia == 2 || dia == 3 || dia == 5 {
		precoFinal = precoNormal * 0.60
	} else {
		precoFinal = precoNormal
	}

	// Aplica acréscimo de 15% se for lançamento
	if categoria == "lançamento" {
		precoFinal = precoFinal*0.15 + precoFinal
	}

	fmt.Printf("Preço final a pagar: R$ %.2f\n", precoFinal)
}
