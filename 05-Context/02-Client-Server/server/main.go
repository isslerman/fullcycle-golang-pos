// Creating context that knows if it was cancelled or still running. Checking ctx cancelled by client to stop some processing.
package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// get the context requisition
	ctx := r.Context()
	log.Println("Request started.")
	defer log.Println("Request ended.")
	select {
	case <-time.After(5 * time.Second):
		// Server log
		log.Println("Request processed with success.")
		// Imprime no browser
		w.Write([]byte("Request processed with success."))
		return
	case <-ctx.Done():
		log.Println("Request cancelled by client.")
	}
}
