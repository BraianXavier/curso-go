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
	//struct para json

	p1 := produto{1, "notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1json, _ := json.Marshal(p1)
	fmt.Println(string(p1json))

	//json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.9,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}