package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1) //Marshal retorna dois valores, o primeiro é um json e o segundo seria um erro
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	JsonString := `{"id":2, "nome":"Caneta", "preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(JsonString), &p2)
	fmt.Println(p2.Tags[1])
}
