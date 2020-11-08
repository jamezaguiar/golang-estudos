package main

import "fmt"

type sportive interface {
	turnOnTurbo()
}

type luxurious interface {
	park()
}

type sportiveLuxurious interface {
	sportive
	luxurious
}

type bmw7 struct{}

func (b bmw7) turnOnTurbo() {
	fmt.Println("Turbo On!")
}

func (b bmw7) park() {
	fmt.Println("Parking!")
}

func main() {
	var b sportiveLuxurious = bmw7{}
	b.turnOnTurbo()
	b.park()
}
