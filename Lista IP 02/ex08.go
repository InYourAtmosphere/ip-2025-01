package main

import (
	f "fmt"
)

func main() {

	var numero int

	f.Scan("%d", &numero)

	if numero < 90 && numero > 20 {

		f.Print("O numero está entre 90 e 20\n")

	} else {

		f.Print(("o numero nao esta entre 90 e 20\n"))
	}

}
