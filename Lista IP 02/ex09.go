package main

import (
	f "fmt"
)

func main() {

	var compra float64

	f.Scanf("%f", &compra)

	if compra < 10 {
		lucro := compra / 0.3
		f.Print("lucro %f", lucro)
	} else if compra >= 10 && compra < 30 {
		lucro := compra / 0.5
		f.Print("lucro %f", lucro)
	} else if compra >= 30 && compra < 50 {
		lucro := compra / 0.6
		f.Print("lucro %f", lucro)
	} else if compra >= 50 {
		lucro := compra / 0.7
		f.Print("lucro %f", lucro)
	}
}
