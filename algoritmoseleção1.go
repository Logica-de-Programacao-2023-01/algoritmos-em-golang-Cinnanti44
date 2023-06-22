package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Printf("%d é o maior número.", num1)
	} else {
		fmt.Printf("%d é o maior número.", num2)
	}

}
