package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct para JSON
	p1 := product{1, "Notebook", 4200, []string{"Informática", "Computadores"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// JSON para Struct
	var p2 product
	jsonString := `{"id":2,"name":"Caneta","price":2,"tags":["Papelaria","Escola"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
}
