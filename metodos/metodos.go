package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando usuario %s no banco de dados\n", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func (u usuario) maiorDeIdade() bool {
	// if u.idade >= 18 {
	// 	fmt.Printf("%s é maior de idade\n", u.nome)
	// } else {
	// 	fmt.Printf("%s é menor de idade\n", u.nome)
	// }

	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Nicollas", 18}
	fmt.Println(usuario1)
	
	usuario1.salvar()
	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}