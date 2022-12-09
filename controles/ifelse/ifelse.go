package main

import "fmt"

func main(){
	imprimirResultado(8.0)
	imprimirResultado(6.5)
}

func imprimirResultado(nota float64){
	if nota >= 7.0 {
		fmt.Println("Aprovado com a nota", nota)
	} else {
		fmt.Println("Reprovado com a nota", nota)
	}
}