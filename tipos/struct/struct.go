package main

import "fmt"

func main (){
	var produto1 produto
	produto1 = produto{
		nome: "Lápis",
		preco: 2.00,
		desconto: 0.05,
	}

	fmt.Printf("%s: %.2f\n", produto1.nome, produto1.calculaDesconto())

	produto2 := produto{"Caneta Preta", 3.00, 0.05}
	fmt.Printf("%s: %.2f\n", produto2.nome , produto2.calculaDesconto())
}

type produto struct {
	nome string
	preco float64
	desconto float64
}

func (p produto) calculaDesconto() float64 {
	return p.preco * (1 - p.desconto)
}