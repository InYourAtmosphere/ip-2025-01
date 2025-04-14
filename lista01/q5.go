package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var conta int
	var consumo float64

	reader := bufio.NewReader(os.Stdin) // lipa o buffer

	input, _ := reader.ReadString('\n') // recebe a entrada

	input = strings.TrimSpace(input) // ler até o enter

	partes := strings.Split(input, " ") // separa por espaço

	conta, _ = strconv.Atoi(partes[0]) // seleciona a primeira parte da string int

	consumo, _ = strconv.ParseFloat(partes[1], 64) // seleciona a segunda parte da string float64

	tipo := []rune(partes[2])[0] // seleicona a terceira parte da sting char

	if tipo == 'C' { // conta comercial
		if consumo <= 80 {
			valor := 500
			fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", conta, valor)
		} else {

			valor := 500 + (0.25 * (consumo - 80))
			fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", conta, valor)

		}
	}

	if tipo == 'R' { // conta residencial

		valor := 5 + 0.05*consumo

		fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", conta, valor)

	}

	if tipo == 'I' { // conta industrial
		if consumo <= 100 {
			valor := 800
			fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", conta, valor)
		} else {

			valor := 800 + (0.04 * (consumo - 100))
			fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", conta, valor)

		}

	}

}

/*
Verificar se antrada tem 3 partes

if len(partes) != 3 {
    fmt.Println("Entrada inválida. Use o formato: <conta> <consumo> <tipo>")
    return
}

Verificar se há erros de conversão

conta, err := strconv.Atoi(partes[0])
if err != nil {
    fmt.Println("Número da conta inválido.")
    return
}

consumo, err := strconv.ParseFloat(partes[1], 64)
if err != nil {
    fmt.Println("Consumo inválido.")
    return
}

Validar o tipo de conta

if tipo != 'C' && tipo != 'R' && tipo != 'I' {
    fmt.Println("Tipo de conta inválido. Use C, R ou I.")
    return
}

*/
