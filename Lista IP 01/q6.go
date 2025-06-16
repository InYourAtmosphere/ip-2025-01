package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	var saidas []string

	for i := 0; i < n; i++ {
		var f float64
		fmt.Scan(&f)

		c := 5.0 * (f - 32.0) / 9.0

		// Formata a string e adiciona ao slice
		linha := fmt.Sprintf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", f, c)
		saidas = append(saidas, linha)
	}

	for _, s := range saidas {
		fmt.Println(s)
	}

}
