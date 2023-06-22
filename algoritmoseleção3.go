package main

import "fmt"

func main() {
	var num1 int

	fmt.Print("Digite o numero:")
	fmt.Scanln(&num1)
	if num1%2 == 0 {
		fmt.Printf("%d é um numero par", num1)
	} else {
		fmt.Printf("%d é um numero impar", num1)
	}
}
