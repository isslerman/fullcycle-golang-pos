package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()

		// Alternative
		// atomic.AddUint64(1)

		time.Sleep(100 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes", number)))
	})
	http.ListenAndServe(":3000", nil)
}

// command-line
// apache benchmark - 10k req, with 100 simult.
// ab -n 10000 -c 100 http://localhost:3000/

// How to check about race problems
// go run -race main.go
