package main

import "fmt"

func main() {
	var numero, maior int

	for {
		fmt.Println("Digite um número inteiro (ou zero para sair):")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		if numero > maior {
			maior = numero
		}
	}

	if maior == 0 {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		fmt.Println("O maior número digitado é:", maior)
	}

}
