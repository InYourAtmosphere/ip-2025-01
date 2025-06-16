package main

import (
	f "fmt"
)

func main() {

	var resto []int
	var numero int

	f.Scanln(&numero)

	for numero > 0 {

		resto = append(resto, numero%16)

		numero = numero / 16

	}

	for i := len(resto) - 1; i >= 0; i-- {
		if resto[i] < 10 {
			f.Printf("%2d", resto[i])
		} else {
			f.Printf("%c", 'A'+(resto[i]-10)) // Converte 10-15 em A-F
		}

	}

}
