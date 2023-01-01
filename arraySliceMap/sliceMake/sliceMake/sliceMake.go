package main

import "fmt"

func main() {
	s := make([]int, 10) // Cria um slice e um array interno 
	fmt.Println(s, len(s), cap(s))

	s = make([]int, 10, 20) // 10 tamanho do slice, 20 tamanho do array interno
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}