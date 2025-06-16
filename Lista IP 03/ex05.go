package main

import (
	"fmt"
)

func main() {
	var idadeS, pesoS, alturaS []float64
	var para, contadorIdade, contador10e20 int
	var media float64
	var idade, peso, altura float64
	var contadorPeso int

	for i := 0; i > -10; i++ {
		fmt.Print("Digite a idade: ")
		fmt.Scanln(&idade)
		idadeS = append(idadeS, idade)

		fmt.Print("Digite o peso: ")
		fmt.Scanln(&peso)
		pesoS = append(pesoS, peso)

		fmt.Print("Digite a altura: ")
		fmt.Scanln(&altura)
		alturaS = append(alturaS, altura)

		fmt.Println("continuar digitando?\n 1-sim diferente de 1 não\n")
		fmt.Scanln(&para)
		if para != 1 {
			break
		}
	}

	contadorIdade = 0
	media = 0
	contador10e20 = 0
	contadorPeso = 0

	for i := 0; i < len(idadeS); i++ {

		if idadeS[i] > 50 {
			contadorIdade++
		}

		if idadeS[i] > 10 && idadeS[i] < 20 {
			media = media + alturaS[i]
			contador10e20++
		}

		if pesoS[i] < 40 {
			contadorPeso++
		}
	}

	var printMedia float64
	if contador10e20 > 0 {
		printMedia = media / float64(contador10e20)
	} else {
		printMedia = 0
	}

	var prinPesos float64
	if len(idadeS) > 0 {
		prinPesos = float64(contadorPeso) / float64(len(idadeS))
	} else {
		prinPesos = 0
	}

	fmt.Printf("quantidade de pessoas 50+: %d\n", contadorIdade)
	fmt.Printf("media das alturas das pessoas com idade entre 10 e 20 anos: %.2f\n", printMedia)
	fmt.Printf("porcentagem de pessoas com peso inferior a 40 quilos: %.2f%%\n", prinPesos*100)
}
