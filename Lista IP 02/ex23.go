package main

import (
	f "fmt"
)

func main() {

	var idade float64

	f.Scanf("%d", &idade)

	if idade < 16 {

		f.Printf("não eleitor")
	}
	if idade > 18 && idade < 65 {

		f.Printf("Eleitor obrigatório")
	} else {

		f.Printf("Eleitor facultativo")
	}

}
