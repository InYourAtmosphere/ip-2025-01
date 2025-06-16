package main

import (
	f "fmt"
)

func main() {

	var numero int

	f.Scanf("%d", &numero)

	if numero%5 == 0 {

		f.Printf("%d e divisivel por 5\n", numero)
	} else {
		f.Printf("%d nao e divisivel por 5\n", numero)
	}
}
