package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var serie, loop int

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()

		valores := strings.Fields(input) // separa por espaço

		if len(valores) == 2 {
			serie, _ = strconv.Atoi(valores[0])

			loop, _ = strconv.Atoi(valores[1])

		}
	}
	if serie%2 != 0 {

		f.Printf("O NUMERO NAO E PAR")

	} else {

		for i := 0; i <= loop; i++ {

			f.Print("%d ", serie)
			serie = serie + 2
		}

		f.Print("\n")
	}

}
