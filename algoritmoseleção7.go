package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Digite o salario:")
	fmt.Scanln(&salario)

	var aumento float64
	if salario < 1000 {
		aumento = 0.1
	} else {
		aumento = 0.05
	}

	novoSalario := salario + (salario * aumento)

	fmt.Println("o novo salario é:", novoSalario)

}
