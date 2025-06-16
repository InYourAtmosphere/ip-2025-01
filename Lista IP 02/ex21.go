package main

import (
	"fmt"
)

func main() {
	var idAluno int
	var nota1, nota2, nota3, mediaExercicios, mediaFinal float64
	var conceito string

	// Leitura dos dados do aluno
	fmt.Print("Digite o número de identificação do aluno: ")
	fmt.Scan(&idAluno)

	fmt.Print("Digite a nota 1: ")
	fmt.Scan(&nota1)

	fmt.Print("Digite a nota 2: ")
	fmt.Scan(&nota2)

	fmt.Print("Digite a nota 3: ")
	fmt.Scan(&nota3)

	fmt.Print("Digite a média dos exercícios: ")
	fmt.Scan(&mediaExercicios)

	// Cálculo da média final
	mediaFinal = (nota1 + (nota2 * 2) + (nota3 * 3) + mediaExercicios) / 7

	// Definir o conceito do aluno com base na média final
	if mediaFinal >= 9.0 && mediaFinal <= 10.0 {
		conceito = "A"
	} else if mediaFinal >= 7.5 && mediaFinal < 9.0 {
		conceito = "B"
	} else if mediaFinal >= 6.0 && mediaFinal < 7.5 {
		conceito = "C"
	} else if mediaFinal >= 4.0 && mediaFinal < 6.0 {
		conceito = "D"
	} else {
		conceito = "E"
	}

	// Determinar aprovação ou reprovação
	status := ""
	if conceito == "A" || conceito == "B" || conceito == "C" {
		status = "APROVADO"
	} else {
		status = "REPROVADO"
	}

	// Exibir resultados
	fmt.Printf("\nNúmero do aluno: %d\n", idAluno)
	fmt.Printf("Nota 1: %.2f\n", nota1)
	fmt.Printf("Nota 2: %.2f\n", nota2)
	fmt.Printf("Nota 3: %.2f\n", nota3)
	fmt.Printf("Média dos exercícios: %.2f\n", mediaExercicios)
	fmt.Printf("Média de Aproveitamento: %.2f\n", mediaFinal)
	fmt.Printf("Conceito: %s\n", conceito)
	fmt.Printf("Status: %s\n", status)
}
