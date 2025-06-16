package main

import (
	f "fmt"
	m "math"
)

func main() {

	var n int

	f.Scanln(&n)

	for i := 1; i <= n; i++ {
		reusltado := m.Pow(float64(i), 2)
		f.Printf(" %4.f ", reusltado)
	}
}
