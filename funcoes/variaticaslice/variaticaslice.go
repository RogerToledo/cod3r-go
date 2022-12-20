package main

import "fmt"

func main(){
	aprovados := []string{"Roger", "Ana", "Yasmin", "Rafael"}
	imprimirAprovados(aprovados...)
}

func imprimirAprovados(aprovados ... string) {
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i + 1, aprovado)
	}
}