package main

import "fmt"

func main() {
	var num int

	fmt.Println("Escreva um numero.")
	fmt.Scan(num)

	if num > 0 {
		fmt.Println("O numero é positivo")

	} else if num < 0 {
		fmt.Println("O numero é negativo")

	} else {
		fmt.Println("O numero é nulo")
	}

}