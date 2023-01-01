package matematica

import (
	"fmt"
	"strconv"
)

func Media(numeros []float64) float64 {
	valor := Arit(numeros, "+")
	if float64(len(numeros)) == 0 {
		return 0.0
	}

	media := valor / float64(len(numeros))
	mediaArredondada, error := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	if error != nil {
		fmt.Println(error)
	}

	return mediaArredondada
}

func Arit(numeros []float64, operacao string) float64 {
	valor := 0.0
	if len(numeros) == 0 {
		return valor
	}

	for indice, numero := range numeros {
		switch {
		case operacao == "+":
			valor += numero
		case operacao == "-":
			if indice == 0 {
				valor = numero
			} else {
				valor -= numero
			}
		case operacao == "*":
			if indice == 0 {
				valor = numero
			} else {
				valor *= numero
			}
		case operacao == "/":
			if indice == 0 {
				valor = numero
			} else {
				if numero == 0 {
					break
				}
				valor /= numero
			}
		}
	}

	valorArredondado, _ := strconv.ParseFloat(fmt.Sprintf("%.2f",valor), 64)
	return valorArredondado
}