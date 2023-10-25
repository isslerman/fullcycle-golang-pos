package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Geckosays string `json:"gecko_says"`
}

func main() {
	r, err := CoingeckoPing()
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Geckosays)
}

func CoingeckoPing() (*Response, error) {
	url := "https://api.coingecko.com/api/v3/ping"
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var r Response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
