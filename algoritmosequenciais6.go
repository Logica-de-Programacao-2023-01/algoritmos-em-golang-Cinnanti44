package main

import "fmt"

func main() {
	var salarioAtual float64

	fmt.Print("Digite o salário atual do funcionário: ")
	fmt.Scanln(&salarioAtual)

	aumento := salarioAtual * 0.15

	novoSalario := salarioAtual + aumento

	fmt.Printf("O novo salário do funcionário é R$ %.2f", novoSalario)

}
