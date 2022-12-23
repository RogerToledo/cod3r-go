package main

import "fmt"

func main() {
	var c esportivoLuxuoso = bmw7{}
	fmt.Println("Esportivo Luxuoso")
	c.ligaTurbo()
	c.fazerBaliza()

	fmt.Println()

	fmt.Println("Esportivo")
	var e espotivo = bmw7{}
	e.ligaTurbo()

}

type espotivo interface {
	ligaTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	espotivo
	luxuoso
	// pode adicionar outros
}

type bmw7 struct {}

func (b bmw7) ligaTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}