package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Println("Digite o primeiro número real:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número real:")
	fmt.Scanln(&num2)

	fmt.Println("Digite o terceiro número real:")
	fmt.Scanln(&num3)

	mediaPonderada := (num1*2 + num2*3 + num3*5) / 10

	fmt.Printf("A média ponderada dos três números é = %.2\n", mediaPonderada)

}
