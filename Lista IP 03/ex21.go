package main

import (
	f "fmt"
)

func main() {

	var b, n int
	var resultado int
	f.Printf("base e expoente\n")
	f.Scanln(&b)
	f.Scanln(&n)

	resultado = 1
	for i := 0; i < n; i++ {

		resultado = b * resultado

	}

	f.Printf("%2d\n", resultado)

}
