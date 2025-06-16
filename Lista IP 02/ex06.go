package main

import (
	f "fmt"
)

func main() {

	var A, B int

	f.Scanf("%d %d", &A, &B)

	if A%B == 0 {

		f.Printf("%d é divisivel por %d", A, B)

	} else {
		f.Printf("%d nao e divisivel por %d", A, B)
	}
}
