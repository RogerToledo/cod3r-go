package main

import (
	"cod3r/mat/matematica"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	numerosAleatorios := gerarNumRandon()

	printValue(numerosAleatorios)
	media := matematica.Media(numerosAleatorios)
	fmt.Printf("Média: %v\n", media)

	fmt.Println("\n======================================")
	fmt.Println()

	numerosAleatorios = gerarNumRandon()
	printValue(numerosAleatorios)
	soma := matematica.Arit(numerosAleatorios, "+")
	fmt.Printf("Soma: %v\n", soma)
}

func printValue(numeros []float64) {
	for _, numero := range numeros {
		fmt.Printf("Valor: %v\n", numero)
	}
	fmt.Println("--------------------------------------")
}

func gerarNumRandon() []float64 {
	rand.Seed(time.Now().UnixNano())
	qtd := rand.Intn(10)
	var numeros []float64
	for i := 0; i < qtd; i++ {
		ran, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64() * 10), 64)
		numeros = append(numeros, ran)
	}

	return numeros
}