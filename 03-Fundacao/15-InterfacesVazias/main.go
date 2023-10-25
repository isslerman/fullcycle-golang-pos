package main

import "fmt"

func main() {
	var varInt interface{} = 10
	var varString interface{} = "Essa é uma string"

	showAnyType(varInt)
	showAnyType(varString)
}

func showAnyType(t interface{}) {
	fmt.Printf("O tipo da varável é %T e o valor é '%v'\n", t, t)
}
