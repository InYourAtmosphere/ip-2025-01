package main

import (
	f "fmt"
)

func main() {

	var n, fatorial int

	f.Scanln(&n)

	if n >= 0 {
		for i := 1; i <= n; i++ {
			fatorial = fatorial * i
		}
	} else {
		return
	}

	f.Printf("%d\n", n)

}
