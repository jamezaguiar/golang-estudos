package main

import "fmt"

type car struct {
	name         string
	currentSpeed float64
}

type ferrari struct {
	car   // Campos anônimos
	turbo bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.currentSpeed = 160
	f.turbo = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.name, f.turbo)
}
