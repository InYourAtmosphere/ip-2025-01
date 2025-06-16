package main

import (
	f "fmt"
	"math/big"
)

func main() {

	grao := big.NewInt(1)
	total := big.NewInt(1)
	dobro := big.NewInt(2)

	f.Printf("Casa 1 graõs: 1\n")
	for i := 1; i < 64; i++ {

		grao.Mul(grao, dobro)

		total.Add(total, grao)
		f.Printf("Casa %d grãos %s\n", i+1, grao)
	}
	f.Printf("total %s\n ", total)
}
