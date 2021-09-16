package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // 2

	go func () {
		escrever("Programando")
		waitGroup.Done() // -1
	}()

	go func () {
		escrever("GoLang!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 2 (-1) (-1) = 0
}

func escrever(text string)  {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}