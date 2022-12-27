package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // sem buffer

	go rotina(c)

	fmt.Println(<- c) // operação bloqueante
	fmt.Println("Foi lido")

}

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Só depois que for lido")
	fmt.Println(<- c) // deadlock
	fmt.Println("Fim") // não será exibido por causa do deadlock
}