package main

import "fmt"

func main() {
	n := 1
	
	func1(n)
	fmt.Println(n)

	func2(&n)
	fmt.Println(n)

}

func func1(n int) {
	n++
}

func func2(n *int) {
	*n++
}
