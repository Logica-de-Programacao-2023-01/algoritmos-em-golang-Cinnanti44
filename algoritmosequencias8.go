package main

import "fmt"

func main() {
	var diasTrabalhados int
	var valorDiaria float64

	// Lê o número de dias trabalhados e o valor da diária
	fmt.Print("Digite o número de dias trabalhados: ")
	fmt.Scanln(&diasTrabalhados)
	fmt.Print("Digite o valor da diária: ")
	fmt.Scanln(&valorDiaria)

	// Calcula o salário
	salario := float64(diasTrabalhados) * valorDiaria

	// Exibe o resultado na tela
	fmt.Printf("O funcionário trabalhou %d dias, recebeu R$ %.2f por dia e seu salário é R$ %.2f.", diasTrabalhados, valorDiaria, salario)
}
