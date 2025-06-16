package main

import f "fmt"

func main() {

	var soma float64

	for i := 0; i < 37; i++ {

		cima := (38 - i) * (37 - i)

		soma = soma + float64(cima/(1+i))
	}

	f.Printf(" %.3f \n", soma)
}
