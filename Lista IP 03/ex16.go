package main

import (
	f "fmt" // Pacote para entrada e saída de dados
)

func main() {
	var termo1, termo2 int
	var n int

	// 1. Ler Entradas
	f.Print("Digite o primeiro termo da série: ")
	f.Scanln(&termo1)

	f.Print("Digite o segundo termo da série: ")
	f.Scanln(&termo2)

	f.Print("Digite o número total de termos (N >= 3): ")
	f.Scanln(&n)

	f.Println("\nSérie de Fetuccine:")

	// 3. Imprimir os Dois Primeiros Termos
	f.Println(termo1) // A0
	f.Println(termo2) // A1

	// 4. Inicializar variáveis de estado para os cálculos futuros
	// aAntigo2 é o termo (i-2), aAntigo1 é o termo (i-1)
	aAntigo2 := termo1
	aAntigo1 := termo2
	var proximoTermo int

	// 5. Loop para Gerar Termos Restantes (a partir do índice 2, que é o terceiro termo)
	// O loop vai de 2 até n-1, pois já imprimimos os termos nos índices 0 e 1.
	for i := 2; i < n; i++ {
		// Verifica se o índice 'i' é par ou ímpar
		if i%2 != 0 { // Se 'i' é ímpar (A3, A5, A7...)
			proximoTermo = aAntigo1 + aAntigo2
		} else { // Se 'i' é par (A2, A4, A6...)
			proximoTermo = aAntigo1 - aAntigo2
		}

		// Imprime o termo recém-calculado
		f.Println(proximoTermo)

		// Atualiza as variáveis para a próxima iteração:
		// O valor que era o termo (i-1) agora se torna o termo (i-2) para o próximo cálculo.
		aAntigo2 = aAntigo1
		// O termo recém-calculado (Ai) agora se torna o termo (i-1) para o próximo cálculo.
		aAntigo1 = proximoTermo
	}
}
