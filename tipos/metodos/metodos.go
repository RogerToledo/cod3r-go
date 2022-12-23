package main

import (
	"fmt"
	"strings"
)

func main() {
	p1 := pessoa{nome: "Roger", sobrenome: "Toledo"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Yasmin Vidal Toledo")
	fmt.Println(p1.getNomeCompleto())
}

type pessoa struct {
	nome string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]

	var sobrenome string
	for i := 1; i < len(partes); i++ {
		if i == 1 {
			sobrenome = partes[i]
			continue
		}
		sobrenome += " " + partes[i]
	}

	p.sobrenome = sobrenome
}