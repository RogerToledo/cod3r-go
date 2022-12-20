package main

import "fmt"

func main(){
	rst, err := fatorial(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rst)
	}
}

func fatorial(n int) (int, error){
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número invalido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		rst := fatorialAnterior * n
		return rst, nil
	}
}