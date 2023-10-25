// testing the Http Client Get with a timeout. If we change from time.second to milisecond we will get a context timeout.
package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{
		Timeout: time.Second,
		// Timeout: time.Microsecond,
	}
	resp, err := c.Get("http://google.com")
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
