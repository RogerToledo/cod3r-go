package main

import "fmt"

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Printf("A Ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}

type carro struct {
	nome string
	velocidadeAtual int
}

type ferrari struct {
	carro // Campos anônimos
	turboLigado bool
}