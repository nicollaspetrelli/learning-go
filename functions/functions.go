package main

import "fmt"

func main() {
	number := somar(1, 2)
	println(number)

	var f = func (txt string) string {
		println(txt)
		return txt
	}

	resultado := f("Função F")
	fmt.Println(resultado)

	resultadoSoma, resultadoSub := calculosMatematicos(10, 14)
	println(resultadoSoma, resultadoSub)

	resultadoSoma2, _ := calculosMatematicos(10, 14)
	println(resultadoSoma2)
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func somar(n1 int, n2 int) int {
	return n1 + n2
}