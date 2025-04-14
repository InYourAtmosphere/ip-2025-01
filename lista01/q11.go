package main

import (
	f "fmt"
)

func main() {

	var avaliacao int

	f.Scanf("%d", &avaliacao)

	if avaliacao%3 == 0 && avaliacao%5 == 0 {
		f.Print("O NUMERO E DIVISIVEL\n")
	} else {
		f.Printf("O NUMERO NAO E DIVISIVEL\n")
	}

}
