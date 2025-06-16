package main

import (
	f "fmt"
)

func main() {

	var entrada float64

	f.Scanf("%f", &entrada)
	if entrada == 2 {
		return
	}
	y := 8 / (2 - entrada)

	f.Printf("%f", y)
}
