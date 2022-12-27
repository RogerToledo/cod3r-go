package main

import (
	"fmt"

	"github.com/cod3r/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func multiplexar(ent1, ent2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(ent1, c)
	go encaminhar(ent2, c)

	return c
}

func main() {
	c := multiplexar(
		html.Titulo("https://www.cod3r.com.br", "http://www.google.com"),
		html.Titulo("https://www.duckduckgo.com", "http://www.youtube.com"),
	)

	fmt.Println(<-c + " | " + <-c)
	fmt.Println(<-c + " | " + <-c)
}