package main

import (
	"fmt"
	"strconv"
)

func main(){
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado
	v := 97 // --> Não converte o 97 em string
	fmt.Println("Teste " + string(v))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(v))
	fmt.Println("Teste " + fmt.Sprint(v))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("True") // "true" ou "1" = true | "false" ou "0" = false
	if b {
		fmt.Println("Verdadeiro")
	}
}