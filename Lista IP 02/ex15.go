package main

import (
	f "fmt"
)

func main() {

	var data int

	f.Scanf("%d", &data)

	ano := data % 10000

	mes := (data / 10000) % 100

	dia := data / 1000000

	f.Printf("%d %d %d", dia, mes, ano)
}
