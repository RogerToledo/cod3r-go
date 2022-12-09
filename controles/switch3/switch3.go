package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(tipo(2))
	fmt.Println(tipo(3.8))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func () {}))
	fmt.Println(tipo(time.Now()))
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}