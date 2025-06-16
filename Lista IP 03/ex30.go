package main

import (
	f "fmt"
	m "math"
)

func main() {

	var volume float64

	for i := 0.0; i <= 20; i = i + 0.5 {

		volume = (m.Pi * m.Pow(i, 3) * 3) / 4
		f.Printf("R = %3.1f v= %4.4f  \n", i, volume)
	}

}
