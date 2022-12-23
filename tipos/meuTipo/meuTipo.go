package main

import "fmt"

func main() {
	notaParaConceito(10.0)
	notaParaConceito(7.5)
	notaParaConceito(6.99)
	notaParaConceito(3.0)
	notaParaConceito(2.1)
}

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim 
}

func notaParaConceito(n nota) {
	if n.entre(9.0, 10.0) {
		fmt.Println("A")
	} else if n.entre(7.0, 8.99) {
		fmt.Println("B")
	} else if n.entre(5.0, 6.99) {
		fmt.Println("C")
	} else if n.entre(3.0, 4.99) {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}