package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array e Slices")

	// Array
	var array[5]string
	array[0] = "Posição 1"
	fmt.Println(array)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)

	// Slice é uma estrutura baseada no array só que melhor =)
	slice := []int{1, 2, 3, 4, 5}

	slice = append(slice, 18) // adiciona um novo elemento no final do slice
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// provando a diferença entre slice e array
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
}