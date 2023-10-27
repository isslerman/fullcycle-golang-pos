package main

import (
	"github.com/isslerman/goexpert/11-Events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello boss", "amq.direct")
}
