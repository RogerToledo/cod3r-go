package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[12345678928] = "Roger"
	aprovados[98765432100] = "Ana"
	aprovados[11111111190] = "Yasmin"
	aprovados[22222222209] = "Kyara"
	aprovados[33333333310] = "Teste"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados{
		fmt.Printf("cpf(%d): %s\n", cpf, nome)
	}

	fmt.Println(aprovados[33333333310])

	delete(aprovados, 33333333310)

	fmt.Println(aprovados[33333333310])

	for cpf := range aprovados {		
		fmt.Printf("cpf(%d)\n", cpf)
	}

	for _, nome := range aprovados {		
		fmt.Printf("nome:%s\n", nome)
	}
} 