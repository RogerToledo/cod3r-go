package main

import "fmt"

func main () {
	var coisa interface{}
	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)
}

type curso struct{
	nome string
}

