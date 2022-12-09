package main

import "fmt"

func main(){
	definirConceito(9.5)
	definirConceito(8.0)
	definirConceito(6.0)
	definirConceito(4.0)
	definirConceito(1.0)

}

func definirConceito(nota float64){
	if nota >= 9 && nota <= 10 {
		fmt.Println("A")
	} else if nota >= 8 && nota < 9 {
		fmt.Println("B")
	} else if nota >= 5 && nota < 8 {
		fmt.Println("C")
	} else if nota >= 3 && nota < 5 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}