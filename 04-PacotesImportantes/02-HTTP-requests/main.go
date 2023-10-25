package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	// o defer atrasa o fechamento. Para não esquecer, ele espera executar tudo para depois fechar o arquivo.
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	// imprime em bytes
	println(res)
	// imprime em string
	println(string(res))
}
