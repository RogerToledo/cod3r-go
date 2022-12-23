package main

import "fmt"

func main() {
	var coisa imprimivel = pessoa{nome: "Roger", sobrenome: "Toledo"}
	fmt.Println(coisa.toString())
	imprimir(coisa)
	imprimir(pessoa{nome: "Ana", sobrenome: "Toledo"})

	fmt.Println()

	imprimir(produto{nome: "Mouse", preco: 15.00})
}

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("O produto %s custa R$ %.2f\n", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
