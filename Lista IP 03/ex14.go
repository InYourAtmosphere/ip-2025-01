package main

import (
	f "fmt"
)

// isPrime verifica se um número é primo
func isPrime(num int) bool {
	if num <= 1 {
		return false // Números menores ou iguais a 1 não são primos
	}
	if num <= 3 {
		return true // 2 e 3 são primos
	}
	if num%2 == 0 || num%3 == 0 {
		return false // Múltiplos de 2 ou 3 não são primos (exceto o próprio 2 e 3)
	}

	// Verifica divisores a partir de 5, pulando de 6 em 6 (otimização)
	// pois todo primo > 3 pode ser escrito na forma 6k ± 1
	// Podemos ir até a raiz quadrada do número para otimização
	for i := 5; float64(i)*float64(i) <= float64(num); i = i + 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false // Se for divisível, não é primo
		}
	}
	return true // Se nenhum divisor foi encontrado, é primo
}

func main() {
	var n1, n2 int

	f.Printf("Digite o primeiro número (N1): ")
	f.Scanln(&n1) // Lê N1

	f.Printf("Digite o segundo número (N2): ")
	f.Scanln(&n2) // Lê N2

	// Garante que N1 seja menor ou igual a N2 para o loop
	if n1 > n2 {
		n1, n2 = n2, n1 // Troca os valores se N1 for maior que N2
	}

	// Cria um slice para armazenar os números primos encontrados
	primosEncontrados := []int{}

	// Itera por todos os números no intervalo [N1, N2]
	for i := n1; i <= n2; i++ {
		if isPrime(i) { // Chama a função auxiliar para verificar se 'i' é primo
			primosEncontrados = append(primosEncontrados, i) // Adiciona o primo ao slice
		}
	}

	// Imprime os números primos encontrados
	if len(primosEncontrados) == 0 {
		f.Println("Nenhum número primo encontrado no intervalo.")
	} else {
		f.Printf("Números primos entre %d e %d:\n", n1, n2)
		for _, primo := range primosEncontrados {
			f.Printf(" %4d ", primo) // Imprime cada primo em uma nova linha
		}
	}
}
