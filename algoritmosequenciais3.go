package main

import "fmt"

func main() {
	var peso, altura float64

	fmt.Println("Digite o seu peso em quilogramas:")
	fmt.Scanln(&peso)

	fmt.Println("Digite a sua altura em metros:")
	fmt.Scanln(&altura)

	imc := peso / (altura * altura)

	fmt.Printf("Seu IMC é: %.2f\n", imc)
}
