package main

import "fmt"

func gradeForConcept(grade float64) string {
	switch int(grade) {
	case 10:
		fallthrough
		/*
			Ao contrário de outras linguagens, o switch case do Go
			para de executar depois de entrar em um caso.
			Exceto se usarmos a palavra reservada acima, fallthrough.
			Ao entrar no caso 10, a execução continuará para o caso abaixo
		*/
	case 9:
		return "A"
	case 8, 7: // Podemos ter um case que abrange mais de um valor
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid grade."
	}
}

func main() {
	fmt.Println(gradeForConcept(10))
}
