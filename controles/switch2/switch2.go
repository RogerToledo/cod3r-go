package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Print("Bom dia!")
	case t.Hour() < 18:
		fmt.Print("Boa tarde!")
	default:
		fmt.Print("Boa noite!")
	}

}

