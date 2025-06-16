package main

import (
	f "fmt"
)

func main() {

	var avaliar int

	f.Scanf("%d", &avaliar)

	if avaliar%2 == 0 {

		f.Printf("%d e par \n", avaliar)
	} else {

		f.Printf("%d e impar \n", avaliar)
	}

}
