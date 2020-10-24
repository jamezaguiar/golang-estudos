package main

import (
	"fmt"
	"time"
)

func WhichType(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println(WhichType(2.3))
	fmt.Println(WhichType(8))
	fmt.Println(WhichType("Hi"))
	fmt.Println(WhichType(func() {}))
	fmt.Println(WhichType(time.Now()))
}
