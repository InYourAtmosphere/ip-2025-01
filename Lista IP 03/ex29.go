package main

import (
	f "fmt"
)

func main() {

	var n, contador int

	f.Scanln(&n)

	for i := 1; i <= n; i++ {

		f.Printf(" %3d ", i)
		contador++
	}

	soma := ((1 + n) * contador) / 2

	f.Printf("\n %3d \n", soma)
}
