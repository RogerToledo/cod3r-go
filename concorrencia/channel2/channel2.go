package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")
	
	a, b := <- c, <- c
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<- c)
}

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second * 3)
	c <- 4 * base
}