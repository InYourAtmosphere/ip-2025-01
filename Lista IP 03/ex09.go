package main

import (
	f "fmt"
)

func main() {

	var nota1, nota2 []float64
	var n, contadorAprovados, contadorExame, contadorReprovados int
	var media, meidaClasse float64
	var soma float64
	soma = 0
	f.Printf("Quantos alunos?\n")
	f.Scanln(&n)
	nota1 = make([]float64, n)
	nota2 = make([]float64, n)
	for i := 0; i < n; i++ {
		f.Scanln(&nota1[i])
		f.Scanln(&nota2[i])
	}

	contadorAprovados = 0
	contadorExame = 0
	contadorReprovados = 0

	for i := 0; i < n; i++ {
		media = (nota1[i] + nota2[i]) / 2
		f.Printf("aluno %d media %f ", i+1, media)
		if media < 3 {
			f.Printf("Reprovado\n")
			contadorAprovados++
		} else if media > 3 && media < 7 {
			f.Printf("Exame\n")
			contadorExame++
		} else {
			f.Printf("aprovado\n")
			contadorReprovados++
		}
		meidaClasse = meidaClasse + media
	}

	soma = meidaClasse / float64(n)

	f.Printf("Aprovados %d, Exame %d, Reprovados %d, media classe %f\n", contadorAprovados, contadorExame, contadorReprovados, soma)
}
