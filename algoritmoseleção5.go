package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("o numero é multiplo de 3 e de 5 ao mesmo tempo")
	} else {
		fmt.Println("o numero não é multiplo de 3 e de 5 ao mesmo tempo ")
	}
}
