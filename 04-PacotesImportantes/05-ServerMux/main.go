package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Boss!"))
	// })

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My blog"})
	http.ListenAndServe(":8080", mux)

	// Podemos criar um segundo mux server rodando no meu app.
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", HomeHandler2)
	http.ListenAndServe(":8090", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Boss!"))
}

func HomeHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Boss 2!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
