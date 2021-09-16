package main

import "fmt"

// passando por clone
func invertSignal(numero int) int {
	return numero * -1
}

// passando por REFERENCIA
func inverterSinalComPonteiro(numero *int)  {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := invertSignal(numero)

	fmt.Println(numeroInvertido);
	fmt.Println(numero);

	novoNumero := 40
	fmt.Println(novoNumero);
	inverterSinalComPonteiro(&novoNumero) // passando por REFERENCIA
	fmt.Println(novoNumero);

}