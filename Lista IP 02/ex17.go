package main

import (
	f "fmt"
)

func main() {

	var consumo, conta float64
	var char byte
	f.Scanf("%c %f", &char, consumo)

	if char == 'R' {

		conta = 5 + consumo*0.05
		f.Printf("%f", conta)
	}
	if char == 'C' {

		if consumo <= 80 {
			conta = 500
			f.Printf("%f", conta)
		}
		if consumo > 80 {
			conta = 500 + 0.25*(consumo-80)
			f.Printf("%f", conta)
		}

	}
	if char == 'I' {
		if consumo <= 100 {
			conta = 800
			f.Printf("%f", conta)
		}
		if consumo > 100 {
			conta = 800 + (consumo-100)*0.04
			f.Printf("%f", conta)
		}
	}
}
