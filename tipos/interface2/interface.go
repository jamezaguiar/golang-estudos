package main

import "fmt"

type sportive interface {
	turnOnTurbo()
}

type ferrari struct {
	model        string
	turbo        bool
	currentSpeed uint
}

func (f *ferrari) turnOnTurbo() {
	f.turbo = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnOnTurbo()

	var car2 sportive = &ferrari{"F40", false, 0}
	car2.turnOnTurbo()

	fmt.Println(car1, car2)
}
