package main

import "fmt"

func main() {
	funcsPorLetra := map[string] map[string] float64 {
		"G": {
			"Gabriela Silva": 12345.00,
			"Guga Pereira": 10223.00,
		},
		"J": {
			"José João": 2000.00,
		},
		"P": {
			"Pedro Junior": 1300.00,
		},
	} 

	for letra, funcionarios := range funcsPorLetra {
		fmt.Println(letra, funcionarios)
	}

	delete(funcsPorLetra, "P")

	fmt.Println()

	for letra, funcionarios := range funcsPorLetra {
		for nome, salario := range funcionarios {
			fmt.Println(letra, nome, salario)
		}
	}


}