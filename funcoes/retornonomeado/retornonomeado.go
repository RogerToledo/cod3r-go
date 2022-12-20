package main

import "fmt"

func main () {
	x, y := troca(1, 2)
	fmt.Println(x, y)

	a, b := troca(8, 9)
	fmt.Println(a, b)
}

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}