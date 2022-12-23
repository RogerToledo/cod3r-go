package main

import "fmt"

func main() {
	var pedido1 pedido
	var item1, item2, item3 item

	item1 = item {
		1,
		5,
		2.00,
	}
	
	item2 = item {
		10,
		20,
		5.00,
	}

	item3 = item {
		7,
		50,
		0.5,
	}

	var itens1 []item
	itens1 = append(itens1, item1)
	itens1 = append(itens1, item2)
	itens1 = append(itens1, item3)

	pedido1 = pedido{
		pedidoId: 1,
		userId: 1,
		itens: itens1,
	}

	valorTotal1 := pedido1.calculaPreco();
	pedido1.imprimePedido(valorTotal1)

	pedido0 := pedido{
		pedidoId: 2,
		userId: 2,
		itens: []item{
			{produtoId: 2, qtde: 4, preco: 10.50},
			{produtoId: 3, qtde: 20, preco: 1.80},
			{produtoId: 4, qtde: 5, preco: 5.50},
			{produtoId: 5, qtde: 8, preco: 2.00},
			{produtoId: 6, qtde: 5, preco: 3.50},
		},
	}

	valorTotal2 := pedido0.calculaPreco()
	pedido0.imprimePedido(valorTotal2)
}

type item struct {
	produtoId int
	qtde int
	preco float64
}

type pedido struct {
	pedidoId int
	userId int
	itens []item
}

func (p pedido) calculaPreco() float64 {
	var valorTotal float64
	for _, item := range p.itens {
		valorTotal += item.preco * float64(item.qtde)
	}
	return valorTotal
}

func (p pedido) imprimePedido(valorTotal float64) {
	fmt.Printf("Pedido: %d\n", p.pedidoId)
	fmt.Println("----------------------------------")
	for _, item := range p.itens {
		fmt.Printf("Item: %d, Preço: %.2f, Qtd: %d\n",item.produtoId, item.preco, item.qtde)
	}
	fmt.Println("__________________________________")
	fmt.Printf("Valor Total: %.2f\n\n", valorTotal)
}