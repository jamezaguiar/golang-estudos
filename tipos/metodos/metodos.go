package main

import (
	"fmt"
	"strings"
)

type user struct {
	firstName string
	lastName  string
}

func (u user) getFullName() string {
	return u.firstName + " " + u.lastName
}

/*
	Usamos o * para passar uma referência, e não uma cópia,
	fazendo com que user seja realmente alterado.
*/
func (u *user) setFullName(fullName string) {
	partials := strings.Split(fullName, " ")
	u.firstName = partials[0]
	u.lastName = partials[1]
}

func main() {
	p1 := user{firstName: "John", lastName: "Doe"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Jamerson Aguiar")
	fmt.Println(p1.getFullName())
}
