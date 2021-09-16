package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(text string) <-chan string  {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}