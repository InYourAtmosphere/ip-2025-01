package main

import (
	f "fmt"
)

func main() {

	var n1, n2, multi int

	f.Scanln(&n1)
	f.Scanln(&n2)

	for i := 1; i <= n2; i++ {

		multi = multi + n1
	}

	f.Printf("%d\n", multi)
}
