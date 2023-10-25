// testing the Http Client Post with a json buffer.
package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	// recebe um byte e retorna um buffer.
	jsonVar := bytes.NewBuffer([]byte(`{"name": "wesley"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
