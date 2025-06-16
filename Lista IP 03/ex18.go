package main

import (
	f "fmt"
)

func main() {

	for i := 1; i <= 10; i++ {

		for j := 1; j <= 10; j++ {

			if i == j {
				f.Printf(" %2d %2d ", i, j)
			} else {
				f.Printf("        ")
			}
		}
		f.Printf("\n")
	}

}
