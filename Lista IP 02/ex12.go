package main

import (
	f "fmt"
)

func main() {

	var idade int

	f.Scan("%d", &idade)

	if idade > 0 && idade < 2 {
		f.Printf("Recém-nascido")
	}
	if idade >= 3 && idade < 11 {
		f.Printf("Criança")
	}
	if idade >= 12 && idade < 19 {
		f.Printf("Adolescente")
	}
	if idade >= 20 && idade < 55 {
		f.Printf("Adulto")
	}
	if idade >= 55 {
		f.Printf("idoso")
	}
}
