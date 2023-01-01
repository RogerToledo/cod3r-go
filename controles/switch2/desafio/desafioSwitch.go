package main

import "fmt"

func main() {
	fmt.Println(definirConceito(10.0), definirConceito1(10.0))
	fmt.Println(definirConceito(9.1), definirConceito1(9.1))
	fmt.Println(definirConceito(8.0), definirConceito1(8.0))
	fmt.Println(definirConceito(5.3), definirConceito1(5.3))
	fmt.Println(definirConceito(3.2), definirConceito1(3.2))
	fmt.Println(definirConceito(0.0), definirConceito1(0.0))

}

func definirConceito(nota float64) string {
	var rst string
	switch {
	case nota >= 9:
		rst = "A"
	case nota >= 8:
		rst = "B"
	case nota >= 5:
		rst = "C"
	case nota >= 3:
		rst = "D"
	default:
		rst = "E"
	}

	return rst
}

func definirConceito1(nota float64) string {
	switch {
		case nota >= 9 && nota <= 10:
			return "A"
		case nota >= 8 && nota < 9:
			return "B"
		case nota >= 5 && nota < 8:
			return "C"
		case nota >= 3 && nota < 5:
			return "D"
		default:
			return "E"
	}
}