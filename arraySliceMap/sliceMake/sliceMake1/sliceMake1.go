package main

import "fmt"

func main(){
	a1 := [5]int{0, 1, 2}
	fmt.Println(a1)

	s1 := a1[:2]
	fmt.Println(s1, len(s1), cap(s1))

	s1 = append(s1, 9)
	fmt.Println(a1, len(a1), cap(a1))
	fmt.Println(s1, len(s1), cap(s1))

	s1 = append(s1, 3)
	fmt.Println(a1, len(a1), cap(a1))
	fmt.Println(s1, len(s1), cap(s1))



}