package main

import (
	f "fmt"
)

func main() {

	var n int

	f.Printf("quantos números a serem somados ?\n")
	f.Scanln(&n)

	fibonacci := []int{0, 1}

	for i := 2; i < n; i++ {
		proximo := fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, proximo)
	}

	// Imprime a sequência
	f.Println("Sequência de Fibonacci:")
	for i, val := range fibonacci {
		f.Printf("Termo %d: %d\n", i+1, val)
	}

}
