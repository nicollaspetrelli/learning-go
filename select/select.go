package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func ()  {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal de msg 1"
		}	
	}()

	go func ()  {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal de msg 2"
		}	
	}()

	for {

		select {
			case messagemCanal1 := <-canal1:
				fmt.Println(messagemCanal1)
			case messagemCanal2 := <-canal2:
				fmt.Println(messagemCanal2)
		}

		// messagemCanal1 := <-canal1
		// fmt.Println(messagemCanal1)
		
		// messagemCanal2 := <-canal2
		// fmt.Println(messagemCanal2)
	}
}