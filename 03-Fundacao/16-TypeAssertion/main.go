package main

import "fmt"

func main() {
	var varInt interface{} = 10
	var varString interface{} = "Essa é uma string"

	println(varInt)
	println(varString)

	println(varInt.(int))
	println(varString.(string))
	// println(varString.(int)) // esse gera erro panic:

	res, err := varString.(int) // aqui não dá panic
	fmt.Printf("O valor da varável é %v e a conversão retornou como erro: %v\n", res, err)

	res2, _ := varString.(int) // aqui não dá panic
	fmt.Printf("O valor da varável é %v \n", res2)

	// Hoje nas versões mais novas, já temos generics, pesquisar.
}
