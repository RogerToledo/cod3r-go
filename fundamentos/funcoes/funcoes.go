package main

import "fmt"

func main() {
	resultado := somar(5, 4)
	imprimir(resultado)
}

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Print(valor)
}