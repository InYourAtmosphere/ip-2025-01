package main

import "fmt"

func main() {
	
	var soma
	var i int

	for i = 1; i <= 100; i++ {
		fmt.Println(i)
	}

	soma= (( 1 + 100 ) * 100)/2 //fórmula de soma de termos de uma PA

	fmt.Println("A soma é", soma)
}