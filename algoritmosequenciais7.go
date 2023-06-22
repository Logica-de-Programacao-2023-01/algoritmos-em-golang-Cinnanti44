package main

import "fmt"

func main() {
	var numero int

	// Lê o número inteiro
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	// Calcula o sucessor e o antecessor
	sucessor := numero + 1
	antecessor := numero - 1

	// Exibe o resultado na tela
	fmt.Printf("O número digitado foi %d, seu antecessor é %d e seu sucessor é %d.", numero, antecessor, sucessor)

}
