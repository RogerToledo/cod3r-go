package main

import "fmt"

func main(){
	numeros := [...]int{1, 2, 3, 4, 5} // Compilador e atribui tamanho do array

	for i, num := range numeros {
		fmt.Printf("Indice: %d => Valor: %d \n", i, num)
	}

	for _, num := range numeros {
		fmt.Printf("Valor: %d \n", num)
	}
}