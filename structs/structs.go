package main

import "fmt"

type usuario struct {
	nome string
	idade int
	email string
	logradouro address
}

type address struct {
	cep string
	rua string
	numero int
}

func main() {
	fmt.Println("Arquivo structs")

	var user usuario
	user.nome = "Nicollas"
	user.email = "contato@nicollas.dev"
	user.idade = 25

	fmt.Println(user)

	endereco := address{
		cep: "12345-678",
		rua: "Rua do teste",
		numero: 123,
	}

	user2 := usuario{"Nicollas", 25, "contato@nicollas.dev", endereco}
	fmt.Println(user2)

	user3 := usuario{
		nome: "Nicollas",
		email: "contato@nicollas.dev",
		logradouro: endereco,
	}
	fmt.Println(user3)

}