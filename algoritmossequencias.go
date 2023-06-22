package main

import "fmt"

func main() {
	var x, y, z, soma int

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&x)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&y)

	fmt.Println("Digite o terceiro número:")
	fmt.Scanln(&z)

	soma = x + y + z

	fmt.Println("A soma dos três números é:", soma)
}
