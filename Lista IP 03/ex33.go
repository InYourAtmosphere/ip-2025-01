package main

import (
	f "fmt"
)

func main() {

	var n1, n2, contador int

	f.Scanln(&n1)
	f.Scanln(&n2)

	for i := 0; n2 <= n1; i++ {

		n1 = n1 - n2

		contador++
	}

	f.Printf("resto: %d quociente %d", n1, contador)
}
