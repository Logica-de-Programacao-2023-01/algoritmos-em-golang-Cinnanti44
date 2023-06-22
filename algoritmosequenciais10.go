package main

import "fmt"

func main() {
	var pesoEmKilos float64

	// Lê o peso em quilos
	fmt.Print("Digite o peso em quilos: ")
	fmt.Scanln(&pesoEmKilos)

	// Converte o peso para libras
	pesoEmLibras := pesoEmKilos * 2.20462

	// Exibe o resultado na tela
	fmt.Printf("O peso em libras é %.2f.", pesoEmLibras)

}
