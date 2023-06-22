package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Print("digite o primeiro numero:")
	fmt.Scanln(&num1)

	fmt.Print("digite o segundo numero:")
	fmt.Scanln(&num2)

	fmt.Print("digite o terceiro numero:")
	fmt.Scanln(&num3)

	if num1 <= num2 && num1 <= num3 {
		fmt.Printf("%d é o menor número.", num1)
	} else if num2 <= num1 && num2 <= num3 {
		fmt.Printf("%d é o menor número.", num2)
	} else {
		fmt.Printf("%d é o menor número.", num3)
	}
}
