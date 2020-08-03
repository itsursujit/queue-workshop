package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"verify-queue/services"
)

func main() {
	// Connect to a nats server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Printf("Error connecting to nats-io: %s\n", err.Error())
	}
	// Starts an encoded conn to transfer json objects
	encodedConnection, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	consumerService := services.NewConsumerService(nc, encodedConnection)
	stopChannel := make(chan int)

	consumerService.ConsumeAsyncMessages(stopChannel)
	consumerService.ConsumeEncodedMessages(stopChannel)
	select {
	case stop := <-stopChannel:
		fmt.Println("Time to rest", stop)
	}
}