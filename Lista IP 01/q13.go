package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
)

func main() {

	var nota float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nota, _ = strconv.ParseFloat(scanner.Text(), 64)

	if nota >= 9 && nota <= 10 {
		f.Printf("NOTA = %.1f CONCEITO = A\n", nota)

	}

	if nota >= 7.5 && nota < 9 {
		f.Printf("NOTA = %.1f CONCEITO = B\n", nota)

	}

	if nota >= 6 && nota < 7.5 {
		f.Printf("NOTA = %.1f CONCEITO = C\n", nota)
	}

	if nota >= 0 && nota < 6 {
		f.Printf("NOTA = %.1f CONCEITO = D\n", nota)

	}

}
