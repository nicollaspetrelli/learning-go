package main

import "fmt"

// RETORNOS NOMEADOS
func foo(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return soma, subtracao
}

// RETORNO NÃO NOMEADO
func bar(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma, sub := foo(10, 20)
	fmt.Println(soma, sub);
}
