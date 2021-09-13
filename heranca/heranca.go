package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
	altura float32
}

type estudante struct {
	pessoa // "herda" os atributos e métodos da pessoa
	curso string
	faculdade string
}

// PSEUDO "HERANÇA"

func main() {
	fmt.Println("Iniciando a aplicação")

	p1 := pessoa{"Nicollas", "Feitosa", 20, 1.80}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia de Software", "FATEC"}
	fmt.Println(e1)
	fmt.Println(e1.nome) // e1.pessoa.nome
}