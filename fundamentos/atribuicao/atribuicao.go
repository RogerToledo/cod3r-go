package main

import "fmt"

func main(){
	var a byte = 3
	fmt.Println(a)

	i := 3
	fmt.Println("i := 3 =", i)
	i += 3
	fmt.Println("i += 3 =", i)
	i -= 3
	fmt.Println("i -= 3 =", i)
	i /= 2
	fmt.Println("i /= 2 =", i)
	i *= 2
	fmt.Println("i *= 2 =", i)
	i %= 2
	fmt.Println("i %= 2 =", i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}