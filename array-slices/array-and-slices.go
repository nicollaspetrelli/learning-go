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

	// Arrays Internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 2) // quando estourado o valor maximo do slice ele duplica a capacidade
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]int, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // lenght
	fmt.Println(cap(slice4)) // capacidade
}