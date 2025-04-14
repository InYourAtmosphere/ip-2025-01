package main

import (
	f "fmt"
)

func main() {

	var soma, a1, razao, nTermos int

	f.Scanf("%d %d %d", &a1, &razao, &nTermos)

	soma = (a1 + a1 + (nTermos-1)*razao) * nTermos / 2

	f.Printf("%d\n", soma)

}
