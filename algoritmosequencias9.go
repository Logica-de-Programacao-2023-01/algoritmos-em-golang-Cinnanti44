package main

import "fmt"

func main() {

	var preco float64

	// Lê o preço do produto
	fmt.Print("Digite o preço do produto: ")
	fmt.Scanln(&preco)

	// Calcula o preço com desconto
	precoComDesconto := preco * 0.9

	// Exibe o resultado na tela
	fmt.Printf("O preço do produto com desconto é R$ %.2f.", precoComDesconto)
}
