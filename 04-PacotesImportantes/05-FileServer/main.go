package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/boss", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellooo Boss!"))
}
