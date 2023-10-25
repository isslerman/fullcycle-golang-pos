package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// pegando primeiros 0 itens
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	// pegando primeiros 4 itens
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	// cortando os dois primeiros e ficando com o restante. descartando os primeiros
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 256) // aqui a capacidade do array é dobrada e não incrementada em 1.
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
