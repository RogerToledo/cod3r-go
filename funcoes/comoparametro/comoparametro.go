package main

import "fmt"

func main() {
	som := exec(soma, 1, 2)
	fmt.Printf("Soma: %d\n", som)

	sub := exec(subtracao, 1, 2)
	fmt.Printf("Subtração: %d\n", sub)

	mult := exec(multiplicacao, 1, 2)
	fmt.Printf("Multiplicação: %d\n", mult)

	div := exec(divisao, 1, 2)
	fmt.Printf("Divisão %d\n", div)

}

func soma(n1, n2 int) int {
	return n1 + n2
}

func subtracao(n1, n2 int) int {
	return n1 - n2
}

func multiplicacao(n1, n2 int) int {
	return n1 * n2
}

func divisao(n1, n2 int) int {
	return n1 / n2
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}