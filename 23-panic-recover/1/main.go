package main

import "fmt"

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

// Como detectar um panico e trabalhar em cima dele agindo de alguma maneira.

func main() {
	defer func() {
		if r := recover(); r != nil { // pegamos o valor de r e vemos se ele é nil. Se for, o panico foi recuperado e o programa não travou. Podemos seguir.
			if r == "panic1" {
				fmt.Println("panic1 recovered")
			}
			if r == "panic2" {
				fmt.Println("panic2 recovered")
			}
		}
	}()

	panic2()
}
