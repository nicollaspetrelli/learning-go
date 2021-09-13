package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("Escrevendo apartir da mian")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("contato@nicollas.dev")
	fmt.Println(erro)
}