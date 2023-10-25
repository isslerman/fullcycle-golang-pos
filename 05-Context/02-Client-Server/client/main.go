// Here on the client side we create a context too that expires in 3 seconds and the request are not processed or >5 seconds that are processed.
package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

// creating a context with timeout of 5 seconds.
func main() {
	// creating the context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// creating the request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	// executing the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

}
