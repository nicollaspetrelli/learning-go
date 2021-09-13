package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos de inteiros
	//int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128

	var numero int16 = 10000
	// var numero int = 10000  se nao especificado o tipo de int, o compilador tenta inferir o tipo com base na arquitetura do seu computador 32x ou 64x
	println(numero)

	// alias
	// INT32 = RUNE

	// BYTE = UINT8
	var num1 byte = 255
	println(num1)

	var numReal float32 = 1.5
	println(numReal)

	//string
	// strings sempre em aspas duplas

	var nome string = "Nicollas"
	println(nome)

	char := 'n'
	println(char)


	// VALOR 0
	// O VALOR INICIAL DE NUMEROS É 0, STRING É UM VALOR VAZIO E ERROS É NULL 

	var texto string
	println(texto)

	// BOOLEANO

	var boleano1 bool = true
	println(boleano1)


	// ERRO
	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}