package main

import "fmt"

var soma = func(num1, num2 int) int {
	return num1 + num2
}

func main(){
	fmt.Printf("Soma: %d\n", soma(4, 5))

	sub := func(num1, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Sub: %d\n", sub(4, 5))

}