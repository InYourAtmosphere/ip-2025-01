package main

import (
	f "fmt"
)

func main() {

	var numero int

	f.Scanf("%d", &numero)

	if numero < 100 || numero > 999 {
		f.Printf("numero valido\n")
		return
	}

	dezena := numero / 10

	extrair := dezena % 10

	f.Printf("%d  \n", extrair)

}
