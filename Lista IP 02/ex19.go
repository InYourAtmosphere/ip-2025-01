package main

import (
	f "fmt"
	m "math"
)

func main() {

	var opcao int
	var volumeCone, areaCone, volumeCilindro, areaCilindro, volumeEsfera, areaEsfera float64
	var rCone, hCone, rCilindro, hCilindro, rEsfera float64
	f.Printf("Conte Reto 1\nCilindro 2\nEsfera 3\n")
	f.Scanln(&opcao)

	if opcao == 1 {
		f.Printf("Raio\n")
		f.Scanln(&rCone)
		f.Printf("Altura\n")
		f.Scanln(&hCone)
		volumeCone = m.Pi * m.Pow(rCone, 2) * hCone / 3
		areaCone = rCone*m.Sqrt(m.Pow(hCone, 2)+m.Pow(rCone, 2))*m.Pi + m.Pi*m.Pow(rCone, 2)
		f.Printf("volume Cone %f\nArea lateral total %f\n", volumeCone, areaCone)
	} else if opcao == 2 {
		f.Printf("Raio\n")
		f.Scanln(&rCilindro)
		f.Printf("Altura\n")
		f.Scanln(&hCilindro)
		volumeCilindro = m.Pi * m.Pow(rCilindro, 2) * hCilindro
		areaCilindro = 2*m.Pi*rCilindro*hCilindro + 2*m.Pi*m.Pow(rCilindro, 2)
		f.Printf("volume cilindro %f\nArea lateral total %f\n", volumeCilindro, areaCilindro)

	} else if opcao == 3 {
		f.Printf("Raio\n")
		f.Scanln(&rEsfera)

		volumeEsfera = 4 * m.Pi * m.Pow(rEsfera, 3) / 3
		areaEsfera = 4 * m.Pi * m.Pow(rEsfera, 2)

		f.Printf("volume Esfera %f\nArea lateral total %f\n", volumeEsfera, areaEsfera)
	}
}
