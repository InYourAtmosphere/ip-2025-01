package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Ler o preço do produto
	f.Print("Digite o preço do produto: ")
	scanner.Scan()                                                    // consome a linha inteira
	precoStr := scanner.Text()                                        // inicia uma variável precoStr e atribui o que foi digitado a ela
	preco, err := strconv.ParseFloat(strings.TrimSpace(precoStr), 64) // atribui a variavel preco o valor float
	/* (preco, err) existem duas condições, se strconv.ParseFloat conseguir converter
	o número string que scanner.Scan() recebeu do usuário atribui a variável preco esse valor
	strings.TrimSpace(precoStr, 64) consome todos os espaços e transforma o texto em float
	*/
	if err != nil {
		f.Println("Erro: preço inválido")
		return
	}

	// Ler o código da condição de pagamento
	f.Println("Escolha a condição de pagamento:")
	f.Println("1 - À vista, dinheiro ou cheque (10'%' de desconto)")
	f.Println("2 - À vista, cartão de crédito (5'%' de desconto)")
	f.Println("3 - Em 2 vezes (preço normal)")
	f.Println("4 - Em 3 vezes (preço com 10'%' de juros)")
	f.Print("Digite o código da condição de pagamento: ")
	scanner.Scan()
	codigoStr := scanner.Text()
	codigo, err := strconv.Atoi(strings.TrimSpace(codigoStr))
	if err != nil {
		f.Println("Erro: código inválido")
		return
	}

	var valorFinal float64

	switch codigo {
	case 1:
		valorFinal = preco * 0.90
	case 2:
		valorFinal = preco * 0.95
	case 3:
		valorFinal = preco
	case 4:
		valorFinal = preco * 1.10
	default:
		f.Println("Código de condição de pagamento inválido.")
		return
	}

	f.Printf("Valor final a ser pago: R$ %.2f\n", valorFinal)
}
