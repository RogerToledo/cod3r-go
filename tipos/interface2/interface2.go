package main

import "fmt"

func main () {
carro1 := ferrari{modelo: "F40", turboLigado: false}
fmt.Println(carro1)
carro1.ligarTurbo()
fmt.Println(carro1)

var carro2 esportivo = &ferrari{modelo: "F8", turboLigado: false}
fmt.Println(carro2)
carro2.ligarTurbo()
fmt.Println(carro2)
}

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo string 
	turboLigado bool
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}