package main

import "fmt"

// RECEBENDO DE 1 A N int, ou seja pode receber quantos parametros forem necessarios
func soma(numeros ...int) int {
	fmt.Println(numeros); // slice

	// ITERANDO UM SLICE
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

// COMBINANDO VALORES FIXOS COM VALOR VARIATICOS
// limitação: não é possivel ter mais de um parametro variatico por função
func write(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero);
	}
}

func main()  {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 200)
	fmt.Println(totalSoma);

	write("Olá Mundo", 1, 2, 3, 5, 7, 99, 200)
}