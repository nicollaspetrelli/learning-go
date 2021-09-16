package main

import "fmt"

func funcao01()  {
	fmt.Println("Executando 1")
}

func funcao02()  {
	fmt.Println("Executando 2")
}

func alunoIsAprovado(n1, n2 float32) bool  {
	defer fmt.Println("Media calculado, resultado será retornado");
	
	fmt.Println("Entrando na função para verificar se o aluno esta aprovado");
	
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao01()
	// DEFER = ADIAR
	funcao02()

	fmt.Println(alunoIsAprovado(7, 8));
	
}