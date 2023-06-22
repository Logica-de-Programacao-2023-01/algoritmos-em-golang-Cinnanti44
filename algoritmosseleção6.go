package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Digite o primeiro numero:")
	fmt.Scanln(&num1)

	fmt.Print("Digite o primeiro numero:")
	fmt.Scanln(&num2)
	if num1 > 0 && num2 > 0 {
		fmt.Println(num1 * num2)
	} else {
		{
			fmt.Println(num1 + num2)
		}
	}

}
