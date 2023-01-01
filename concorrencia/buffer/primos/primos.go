package main

import "fmt"

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for j := inicio; ; j++ {
			if isPrimo(j) {
				c <- j
				inicio = j + 1
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	primos(cap(c), c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
}
