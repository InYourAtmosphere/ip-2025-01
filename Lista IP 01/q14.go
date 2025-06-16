package main

import (
	f "fmt"
	"math"
)

func main() {

	var alturaPiramide, arestaHexagono float64

	f.Scanf("%f %f", &alturaPiramide, &arestaHexagono)

	volumePiramide := (((3 * math.Pow(arestaHexagono, 2) * math.Sqrt(3)) / 2) * alturaPiramide) / 3

	f.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", volumePiramide)
}
