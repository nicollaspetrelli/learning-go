package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	// Ponteiros é uma forma de acessar o endereço de memória de um valor

	// problema:

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println("variavel1:", variavel1, "variavel2:", variavel2)

	variavel1++
	fmt.Println("variavel1:", variavel1, "variavel2:", variavel2)

	// solução:

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // referenciando o endereço de memoria

	fmt.Println(variavel3, ponteiro) // 100 0xc0000a0a0 - exibe o endereço de memoria
	fmt.Println(variavel3, *ponteiro) // desreferenciando o ponteiro
}