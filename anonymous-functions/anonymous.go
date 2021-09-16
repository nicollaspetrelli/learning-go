package main

import "fmt"

func main() {
	

	// função anonima void executada logo a seguir
	func ()  {
		fmt.Println("Sou uma função anonima e já sou executada");
	}()

	// função anonimo com parametro
	func (text string)  {
		fmt.Println(text);
	}("Passando parametro")

	// função anonima armazenada o retorno
	retorno := func (text string) string  {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Passando parametro")

	fmt.Println(retorno);
	

}