package main

import "fmt"

func main() {
	var num int
	fmt.Println("Digite um número inteiro:")
	fmt.Scanln(&num)

	fmt.Println("O dobro do número é:", num*2)
	fmt.Println("O triplo do número é:", num*3)
	fmt.Println("O quadruplo do número é:", num*4)

}
