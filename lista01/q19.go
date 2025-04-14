package main

import (
	f "fmt"
)

func main() {

	var somaSerie float64 = 0
	var nTermoSerie int

	f.Scanf("%d", &nTermoSerie)

	if nTermoSerie < 1 {
		f.Print("Numero invalido!")
		return
	}

	for i := 1; i <= nTermoSerie; i++ {

		somaSerie = somaSerie + 1/float64(i)
	}

	f.Printf("%.6f\n", somaSerie)
}
