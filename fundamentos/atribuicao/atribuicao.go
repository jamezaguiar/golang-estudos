package main

import "fmt"

func main() {
	var b uint8 = 3
	fmt.Println(b)

	i := 3
	fmt.Println(i)
	i += 3
	fmt.Println(i)
	i -= 3
	fmt.Println(i)
	i /= 2
	fmt.Println(i)
	i *= 2
	fmt.Println(i)
	i %= 2
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
