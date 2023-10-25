package main

import "fmt"

func somaComPonteiros(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {

	a := 10
	fmt.Printf("Var: a | End: %v | Val: %v \n", &a, a)

	var ponteiro *int = &a
	*ponteiro = 20
	fmt.Printf("Var: a | End: %v | Val: %v \n", &a, a)

	// b é um ponteiro
	b := &a
	// imprime o ponteiro.
	fmt.Printf("Var: b | End: %v | Val: %v \n", &b, b)

	c := 10

	// aqui estamos passando o ponteiro e na função é alterado o valor de a
	somaComPonteiros(&a, &c)
	fmt.Printf("Var: a | End: %v | Val: %v \n", &a, a)

	// nessa função, passamos uma cópia do valor das variáveis
	// soma(a,b)
	// aqui passamos os endereços de memória diretamente aonde estão os valores
	// soma(&a, &b)
}
