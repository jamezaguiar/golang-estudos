package main

import (
	"fmt"
	"time"
)

func speak(person, text string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", person, text, i+1)
	}
}

func main() {
	// Execução em série (Uma chamada após a outra)
	/* 	speak("Mary", "Why you don't speak with me?", 3)
	   	speak("John", "I can only speak after you.", 1) */

	// Go Routines
	go speak("Mary", "Why you don't speak with me?", 3)
	go speak("John", "I can only speak after you.", 1)
}
