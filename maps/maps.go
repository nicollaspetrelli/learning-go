package main

import "fmt"

func main() {
	fmt.Println("Maps!")

	usuario := map[string]string {
		"nome": "Nicollas",
		"sobrenome": "Feitosa",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "Nicollas",
			"ultimo": "Feitosa",
		},
		"curso": {
			"nome": "Go",
		},
	}

	usuario2["datanasc"] = map[string]string {
		"ano": "2019",
	}

	// fmt.Println(usuario2["nome"]["primeiro"])
	// fmt.Println(usuario2["curso"]["nome"])
	
	fmt.Println(usuario2)
	
	delete(usuario2, "curso")
	
	fmt.Println(usuario2)
	
	
}