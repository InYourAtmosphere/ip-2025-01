package main

import (
	f "fmt"
)

func main() {

	var n1, n2, minimo int
	minimo = 1

	f.Scanln(&n1)
	f.Scanln(&n2)

	if n1 > n2 {

		for i := 1; i <= n1; i++ {

			if n1%i == 0 && n2%i == 0 {
				minimo = i * minimo
			}

		}

	} else {

		for i := 1; i <= n2; i++ {

			if n1%i == 0 && n2%i == 0 {
				minimo = i * minimo
			}
		}

	}

	f.Printf("%d\n", minimo)

}
