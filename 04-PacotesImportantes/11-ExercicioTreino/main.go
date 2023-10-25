// Get the unclaimed fees from the page https://app.uniswap.org/pool/809174?chain=arbitrum
package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

const timeout = 30

func main() {
	ctx := context.Background()
	// here he have a context that expires in one second
	ctx, cancel := context.WithTimeout(ctx, time.Second*timeout)
	// here the context will be cancelled only if you run the cancel()
	// ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://app.uniswap.org/pool/809174?chain=arbitrum", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
