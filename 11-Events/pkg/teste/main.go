package main

import "fmt"

func main() {
	eventos := []string{"evento1", "evento2", "evento3", "evento4"}

	eventos2 := eventos[1:]
	eventos = append(eventos[:0], eventos[1:]...)
	fmt.Println(eventos)
	fmt.Println(eventos2)
}
