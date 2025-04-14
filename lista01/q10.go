package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
)

func main() {

	var A, B, C, D, det float64

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	A, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	B, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	C, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	D, _ = strconv.ParseFloat(scanner.Text(), 64)

	det = A*D - B*C

	f.Printf("O VALOR DE DETERMINANTE E = %.2f\n", det)
}
