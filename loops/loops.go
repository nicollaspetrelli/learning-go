package main

import (
	"fmt"
	"time"
)

// No GO só temos o FOR como estrutura de repetição

func main() {
	fmt.Println("Loops")

	// ============ while

	i := 0

	for i < 10 {
		// time.Sleep(time.Second)
		fmt.Println("Incrementado I ");
		i++
		fmt.Println(i);
	}

	fmt.Println("Fim do loop")
	fmt.Println(i);
	

	// ===================== for convencional

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementado J ");
		fmt.Println(j);
	}

	// ===================== range

	nomes := [3]string{"Nicollas", "Kebão", "Morilongo"} // array

	for indice, valor := range nomes {
		fmt.Println(indice, valor);
	}

	for _, nome := range nomes {
		fmt.Println(nome);
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)); // se usar sem o metodo string retorna os numero da tabela ASCII
	}

	usuario := map[string]string {
		"nome": "Nicollas",
		"sobrenome": "Feitosa",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor);
	}


	// loop infinito

	for {
		fmt.Println("Executando infinitamente");
		time.Sleep(time.Second)
	}
}
