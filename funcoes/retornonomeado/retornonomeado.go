package main

import "fmt"

func toggle(fst, snd int) (second, first int) {
	second = snd
	first = fst
	return // Retorno limpo. second e first serão retornamos, como definido anteriormente.
}

func main() {
	x, y := toggle(1, 2)
	fmt.Println(x, y)
}
