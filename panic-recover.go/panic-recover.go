package main

import "fmt"


func recuperarExecucao()  {
	fmt.Println("Tentativa de recuperar a execução!");

	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso!");
	}
}

func alunoIsAprovado(n1, n2 float32) bool  {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6!")
}


func main() {
	fmt.Println(alunoIsAprovado(6, 6));
	fmt.Println("Pos execucao!");
}