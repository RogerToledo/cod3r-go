package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// struct para json
	fmt.Println("struct para json")
	p1 := produto{
		Id: 1,
		Nome: "Notebook",
		Preco: 1899.00,
		Tags: []string{"Promoção", "Eletrônicos"},
	}
	fmt.Println(p1)
	p1ToJson, _ := json.Marshal(p1)
	fmt.Println(p1ToJson)
	fmt.Println(string(p1ToJson))

	fmt.Println()

	// json para struct
	fmt.Println("json para struct")
	var p2 produto
	jsonString := `{"id":2, "nome":"Caneta", "preco":8.90, "tags":["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Tags[1])

	fmt.Println()

	var p1ToStruct produto
	fmt.Println(p1ToJson)
	json.Unmarshal(p1ToJson, &p1ToStruct)
	fmt.Println(p1ToStruct)
	fmt.Println(p1ToStruct.Tags[0])
}

type produto struct {
	Id    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}
