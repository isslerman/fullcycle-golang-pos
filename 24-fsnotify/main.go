package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB   string `json:"db"`
	Host string `json:"host"`
	User string `json:"user"`
	Pass string `json:"password"`
}

var config DBConfig

// Usamos o fsnotify para verificar se o arquivo de configuração foi alterado e carregar ele novamente.
// HOW TO RUN:
// go run main.go
// faça e salve alguma alteração no config.json
func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	MarshalConfig("config.json")
	fmt.Println(config)

	done := make(chan bool)
	go func() {
		for {
			// aqui usamos o case select com o case que ficará aguardando algum evento.
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig("config.json")
					fmt.Println("modified file:", event.Name)
					fmt.Println(config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()
	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}
	<-done
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}
