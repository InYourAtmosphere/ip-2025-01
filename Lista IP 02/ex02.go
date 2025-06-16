package main

import (
	"fmt"
)

func main() {

	var avaliar float64

	fmt.Scanf("%f", &avaliar)

	if avaliar > 0 {
		fmt.Printf("%f  e positivo\n", avaliar)
	} else if avaliar < 0 {
		fmt.Printf("%f  e negativo\n", avaliar)
	} else {
		fmt.Printf("%f e nulo\n", avaliar)
	}
}
