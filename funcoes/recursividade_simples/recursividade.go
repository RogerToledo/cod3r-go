package main

import "fmt"

func main(){
	rst := fatorial(5)
	
	fmt.Println(rst)
}

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		fatorialAnterior := fatorial(n - 1)
		rst := fatorialAnterior * n
		return rst
	}
}