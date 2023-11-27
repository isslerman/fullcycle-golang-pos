package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// middleware para o servidor web. Quando receber um panic, volta o server e retorna o panic.
// temos uma função que analisa o panic recebido e usando a função recover, conseguimos ver se é possivel restaurar e seguir.
// HOW TO RUN:
// go run 2/main.go
// curl localhost:3000/panic
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Boss"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	log.Println("Listening on", ":3000")
	if err := http.ListenAndServe(":3000", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", ":3000", err)
	}
}

// passamos um mux e retornamos o mix
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic: %v\n", r)
				debug.PrintStack()
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		fmt.Println("Entrou no handler")
		next.ServeHTTP(w, r)
	})
}
