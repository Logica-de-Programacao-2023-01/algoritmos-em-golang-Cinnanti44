package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Digite a idade:")
	fmt.Scanln(&idade)
	if idade <= 9 {
		fmt.Println("mirim")

	}
	if idade <= 13 && idade >= 10 {
		fmt.Println("infantil")
	}
	if idade <= 17 && idade >= 14 {
		fmt.Println("juvenil")
	}

	if idade >= 18 {
		fmt.Println("adulto")
	}

}
