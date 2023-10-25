package main

// o uso de mais canais nem sempre melhora o benchmark,
// em geral costuma apenas consumir mais memória sem melhorar performance.
func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "Boss"

	println(<-ch)
	println(<-ch)
}
