package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// Podemos iniciar variáveis em um if
	// A varíavel i estará disponível dentro do escopo do if
	if i := randomNumber(); i > 5 {
		fmt.Println("Ganhou com", i)
	} else {
		fmt.Println("Perdeu com", i)
	}
}
