package main

import (
	f "fmt"
)

func main() {

	var horas, minutos, segundos, tempoSegundos int

	f.Scanf("%d", &horas)
	f.Scanf("%d", &minutos)
	f.Scanf("%d", &segundos)

	tempoSegundos = horas*3600 + minutos*60 + segundos

	f.Printf("O TEMPO EM SEGUNDOS E = %d\n", tempoSegundos)

}
