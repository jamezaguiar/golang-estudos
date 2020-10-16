package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // O tipo da variável será inferida pelo compilador

	area := PI * math.Pow(raio, 2) // Forma reduzida de declarar uma var
	// O programa não será compilado se existirem variáveis não utilizadas
	fmt.Println(area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(c, d)

	// É possivel declarar várias variáveis em apenas uma linha
	var e, f bool = true, false
	fmt.Println(e, f)

	// E também usando a sintaxe reduzida
	g, h, i := 2, false, "hi"
	fmt.Println(g, h, i)
}
