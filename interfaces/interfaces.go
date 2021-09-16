package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma)  {
	fmt.Printf("A area da forma é %0.2f \n", f.area())
}

func main() {
	ret := retangulo{altura: 10, largura: 15}
	escreverArea(ret)

	cir := circulo{raio: 10}
	escreverArea(cir)
}