package main

import (
	"github.com/sandronister/generate-messages/pkg/service"
	"github.com/sandronister/go_broker/pkg/broker/factory"
)

func main() {

	broker := factory.NewBroker(factory.REDIS, "localhost", 6379)

	messages := service.GenerateMessage()

	for _, message := range messages {
		err := broker.Producer(&message)
		if err != nil {
			panic(err)
		}
	}

}
