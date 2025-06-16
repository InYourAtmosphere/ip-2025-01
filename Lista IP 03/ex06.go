package main

import (
	f "fmt"
)

func main() {

	var numero, q int

	f.Scanln(&numero)
	q = 0
	for i := 0; i < numero; i++ {

		if i*(i+1)*(i+2) == numero {
			f.Printf("e triangular\n")
			q++
		}

	}

	if q == 0 {
		f.Printf("não é triangular\n")
	}

}
