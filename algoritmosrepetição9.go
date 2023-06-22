package main

import "fmt"

func main() {
	var numero, soma int
	var quantidade int

	for {
		fmt.Println("Digite um número inteiro (ou zero para sair):")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		soma += numero
		quantidade++
	}

	if quantidade == 0 {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		media := float64(soma) / float64(quantidade)
		fmt.Println("A média aritmética dos números digitados é:", media)
	}

}
