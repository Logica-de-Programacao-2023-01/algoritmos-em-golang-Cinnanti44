package main

import "fmt"

func main() {
	var idadeAnos int

	fmt.Println("Digite a idade da pessoa em anos:")
	fmt.Scanln(&idadeAnos)

	idadeDias := idadeAnos * 365

	fmt.Printf("A idade da pessoa em dias é: %"+
		"d\n ", idadeDias)

}
