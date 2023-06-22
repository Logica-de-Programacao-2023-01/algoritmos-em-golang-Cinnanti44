package main

import "fmt"

func main() {
	var altura float64
	var sexo string

	fmt.Println("Digite a altura da pessoa em metros:")
	fmt.Scanln(&altura)

	fmt.Println("Digite o sexo da pessoa (M para masculino e F para feminino):")
	fmt.Scanln(&sexo)

	var pesoIdealMin, pesoIdealMax float64

	if sexo == "M" {
		pesoIdealMin = 20 * altura * altura
		pesoIdealMax = 25 * altura * altura
	} else if sexo == "F" {
		pesoIdealMin = 19 * altura * altura
		pesoIdealMax = 24 * altura * altura
	} else {
		fmt.Println("Sexo inválido.")
		return
	}

	fmt.Printf("O peso ideal para a altura %.2fm é entre %.2fkg e %.2fkg.\n", altura, pesoIdealMin, pesoIdealMax)

	var peso float64

	fmt.Println("Digite o peso da pessoa em quilogramas:")
	fmt.Scanln(&peso)

	imc := peso / (altura * altura)

	if imc < 18.5 {
		fmt.Println("A pessoa está abaixo do peso ideal.")
	} else if imc < 25 {
		fmt.Println("A pessoa está dentro do peso ideal.")
	} else {
		fmt.Println("A pessoa está acima do peso ideal.")
	}
}
