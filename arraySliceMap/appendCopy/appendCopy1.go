package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//a := [3]int{1, 2, 3}
	n := 5
	s := make([]int, 5)
	fmt.Println(s)

	for i := 0; i <= n; i++ {
		s = append(s, aleatorio())
	}
	fmt.Println(s)

	for i := 0; i < 4; i++ {
		s[i] = aleatorio()
	}
	fmt.Println(s)
	
}

func aleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}