package main

import "fmt"

func main() {
	var num int
	fmt.Print("digite um numero:")
	fmt.Scanln(&num)

	for i := 1; i <= 10; i++ {
		println(i * num)
	}

}
