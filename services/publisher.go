package services

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"verify-queue/models"
)

type PublisherService struct {
	natsConnection *nats.Conn
	encodedConn    *nats.EncodedConn
}

// NewService creates a new service
func NewPublisherService(nc *nats.Conn, ec *nats.EncodedConn) *PublisherService {
	return &PublisherService{natsConnection: nc, encodedConn: ec}
}

//Publish string messages
func (c *PublisherService) PublishMessage(message *models.Message) {
	// Simple Publisher
	err := c.natsConnection.Publish("msgChannel", []byte(message.Message))
	if err != nil {
		fmt.Println(err)
	}
}

//Publish encoded messages
func (c *PublisherService) PublishEncodedMessage(message *models.Message) {
	// Simple Encoded Publisher
	err := c.encodedConn.Publish("encodedChannel", message)
	if err != nil {
		fmt.Println(err)
	}
}