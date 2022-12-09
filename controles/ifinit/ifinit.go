package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	if i := numeroAleatorio(); i >= 5 {
		fmt.Println(i, "- Ganhou!!")
	}else{
		fmt.Println(i, "- Perdeu!")
	}

}

func numeroAleatorio() int{
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10) 
}