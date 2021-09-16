package main

import "fmt"

var n int

// pode ter uma por aquivo
func init()  {
	fmt.Println("Executando a função init")
	n = 10
}

// função main pode ter uma por pacote
func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}