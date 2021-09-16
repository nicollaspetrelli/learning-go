package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola Mundo!", canal)
	
	fmt.Println("Depois da função escrever começar a ser executada");
	

	// Evitando Deadlock

	for msg := range canal {
		fmt.Println(msg);
	}

	// for { 
	// 	msg, isOpen := <- canal // recebendo um valor
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(msg);
	// }

	fmt.Println("Fim do programada!");

	// dead lock
	// for { 
	// 	msg := <- canal // recebendo um valor
	// 	fmt.Println(msg);
	// }
}

func escrever(text string, canal chan string)  {
	for i := 0; i < 5; i++ {
		canal <- text // canal recebe texto
		time.Sleep(time.Second)
	}

	close(canal)
}