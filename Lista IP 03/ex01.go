package main

import (
	f "fmt"
)

func main() {

	var base, expoente float64

	f.Printf("escreva a base e o expoente\n")
	f.Scanf("%f %f", &base, &expoente)

	soma := 1.0 //obs se colocar 0, soma será int, se for 0.0 será float

	for i := 1.0; i <= (expoente); i++ {

		soma = (soma * base)

	}

	f.Printf("resultado: %.2f\n", soma)

}
