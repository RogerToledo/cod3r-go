package main

import "fmt"

func main() {
	rst := media(2.0, 8.4, 5.5, 6.7)
	fmt.Printf("Média: %.2f", rst)
}

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	media := total / float64(len(numeros))

	return media
}