package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ir, inss, salario float64
	var horasTrabalhadas, matricula int

	scanner := bufio.NewScanner(os.Stdin)

	f.Printf("matricula funionario:\n")
	scanner.Scan()
	matriculaStr := scanner.Text()
	matricula, _ = strconv.Atoi(strings.TrimSpace(matriculaStr))

	f.Printf("quantidade de horas extras trabalhadas\n")
	scanner.Scan()
	horasTrabalhadasStr := scanner.Text()
	horasTrabalhadas, _ = strconv.Atoi(strings.TrimSpace(horasTrabalhadasStr))

	salario = float64(3*788 + 10*(horasTrabalhadas))

	if salario > 1500 {

		inss = 0.12 * salario
		ir = 0
	}

	if salario > 2000 {

		inss = 0.12 * salario
		ir = 0.2 * salario
	}

	renda := salario - inss - ir
	matricula = matricula + 0
	f.Printf("salario bruto: %f\nsalario liquido %f", salario, renda)
}
