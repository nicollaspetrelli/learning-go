package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá!"), escrever("Go Lang!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEnrtrada2 <-chan string) <-chan string {
	canalDeSaida :=	make(chan string)

	go func () {
		for {
			select {
				case mensagem := <-canalDeEntrada1:
					canalDeSaida <- mensagem
				case mensagem := <-canalDeEnrtrada2:
					canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(text string) <-chan string  {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
