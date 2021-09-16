package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO

	go escrever("Nicollas Fera!") // GoRoutine
	escrever("Programando em Go")

}

func escrever(text string)  {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}