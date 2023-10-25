package main

import (
	"fmt"

	"curso-go/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v", s)
	fmt.Printf("Resultado: %v", matematica.VarAcessivelExternal)
	// fmt.Printf("Resultado: %v", matematica.varNotAcessivelExternal)

	fmt.Println(uuid.New())
}
