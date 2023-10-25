package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criando arquivo
	f, err := os.OpenFile("arquivo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Escrevendo no arquivo
	// tamanho, err := f.WriteString("Hello Boss!")
	tamanho, err := f.Write([]byte("Hello Boss!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)

	// Leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("\nLendo arquivo:")
	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo
	fmt.Println("\nLendo arquivo 2:")
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
