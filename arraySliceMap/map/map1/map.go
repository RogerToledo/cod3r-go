package main

import "fmt"

func main() {
	funcsSalario := map[string]float64 {
		"Maria" : 10000.00,
		"Jose"  : 11000.00,
		"Rafael": 20000.00,
	}

	for nome, salario := range funcsSalario {
		fmt.Printf(nome + " = " + fmt.Sprint(salario) + "\n")
	}

	funcsSalario["Roger"] = 25000.00
	fmt.Println("------------------------------------------")

	for nome, salario := range funcsSalario {
		fmt.Printf(nome + " = " + fmt.Sprint(salario) + "\n")
	}

	delete(funcsSalario, "Ana")

	funcsSalario["Rafael"] = 25000.00

	fmt.Println("------------------------------------------")

	for nome, salario := range funcsSalario {
		fmt.Printf(nome + " = " + fmt.Sprint(salario) + "\n")
	}

	fmt.Println("------------------------------------------")

	for nome := range funcsSalario {
		fmt.Println(nome)
	}

	fmt.Println("------------------------------------------")

	for _, salario := range funcsSalario {
		fmt.Println(salario)
	}

}