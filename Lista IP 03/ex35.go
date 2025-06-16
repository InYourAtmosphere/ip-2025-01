package main

import (
	f "fmt"
)

func main() {

	var resto []int
	var numero int

	f.Scanln(&numero)

	for numero > 0 {

		resto = append(resto, numero%2)

		numero = numero / 2

	}

	for i := len(resto) - 1; i >= 0; i-- {
		f.Printf("%2d", resto[i])
	}

}
