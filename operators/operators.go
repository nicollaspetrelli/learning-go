package main

import "fmt"

func main() {
	// ARITIMETIC OPERATORS
	soma := 1 + 2
	sub := 1 - 2
	mult := 1 * 2
	div := 1 / 2
	mod := 1 % 2

	fmt.Println(soma, sub, mult, div, mod)

	var num1 int16 = 10
	var num2 int16 = 15

	var soma2 int16 = num1 + num2
	fmt.Println(soma2)

	// OPERADORES ATRIBUIÇÃO

	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2) // maior ou menos
	fmt.Println(1 >= 2) // maior/menor ou igual
	fmt.Println(1 == 2) // comparação
	fmt.Println(1 != 2)

	// OPERADORES LOGICOS
	fmt.Println(true && true) // AND
	fmt.Println(true && false) 
	fmt.Println(true || true) // OR
	fmt.Println(true || false)
	fmt.Println(!true) // negação

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero += 10
	fmt.Println(numero)
	numero -= 10
	fmt.Println(numero) 
	numero *= 10 // numero = numero * 10
	fmt.Println(numero)
	numero /= 10 // numero = numero / 10
	fmt.Println(numero)
	numero %= 10
	fmt.Println(numero)

	// OPERADOR TERNARIO NAO FUNCIONA COM GO

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

	// OPERADORES DE PONTEIRO
	variavel3 := "String3"
	fmt.Println(&variavel3)
	fmt.Println(*&variavel3) // * é o ponteiro



}