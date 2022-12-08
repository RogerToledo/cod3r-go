package main

import "fmt"

func main(){
	i := 1

	var p *int = nil
	p = &i // pegando endereço da variável
	*p++   // desreferenciando (pegando valor)
	i++

	fmt.Println(p, *p, &i, i)
	
	// Go não tem aritimética de ponteiro
	// p++ - tem que usar valor *p

	
}