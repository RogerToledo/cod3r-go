package main

import "fmt"

func main(){
	fmt.Print("Mesma")
	fmt.Print(" Linha!")
	
	fmt.Print("\n")
	
	fmt.Println("Pula")
	fmt.Println("Linha!")

	x := 3.141516

	xs := fmt.Sprintln(x)

	fmt.Println("O valor de x é " + xs)
	fmt.Println(">> Outra forma >> x =", x)
	fmt.Printf(">> Outra forma >> x = %.2f", x)
	fmt.Printf("\n>> Outra forma >> x = %.3f", x)
	fmt.Println("\n>> Outra forma >> x =", x, "!!!")

	a := 1
	b := 1.4444
	c := false
	d := "opa"

	fmt.Printf("\n %d | %.2f | %t | %s", a, b, c, d)
	fmt.Printf("\n %v | %v | %v | %v", a, b, c, d)
	fmt.Printf("\n %v | %.2v | %v | %v", a, b, c, d)
}