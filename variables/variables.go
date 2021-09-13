package main

import "fmt"

func main() {
	var name string = "Nicollas"
	secoundName := "Feitosa" // inferencia de tipo (tipo inplicito)

	fmt.Println(name, secoundName)

	var (
		var1 string = "var1"
		var2 string = "var2"
	)
	
	fmt.Println(var1, var2)


	var3, var4 := "var3", "var4";

	fmt.Println(var3, var4)

	const constante1 string = "Const 1"
	fmt.Println(constante1)

	var3, var4 = var4, var3;
	fmt.Println(var3, var4)
}