package main

import (
	"encoding/json"
	"os"
)

type Token struct {
	Name    string `json:"nome"`
	Code    string `json:"codigo"`
	Network string `json:"-"`
}

func main() {
	token := Token{Name: "Solana", Code: "SOL"}

	// 1 - ENCODE
	/////////////

	// retorna o json na variavel
	res, err := json.Marshal(token)
	if err != nil {
		print(err)
	}
	println(string(res))

	// método 1
	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(token)
	// método 2 - short
	err = json.NewEncoder(os.Stdout).Encode(token)
	if err != nil {
		print(err)
	}

	// 2 - DECODE
	/////////////
	jsonPuro := []byte(`{"nome":"DAI", "codigo":"USDC"}`)
	var tokenB Token
	err = json.Unmarshal(jsonPuro, &tokenB)
	if err != nil {
		print(err)
	}
	println(tokenB.Name)
}
