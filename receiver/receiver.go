package main

import (
	"fmt"
	"log"

	connection "farukh.go/micro/connection"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, ch := connection.DeclareConnectionAndCreateChannel()
	defer ch.Close()
	defer conn.Close()
	qName := connection.DeclareQWithBinding(ch)
	messageChannel := getReceiveChannel(ch, qName)
	for delivery := range messageChannel {
		fmt.Printf("%s", delivery.Body)
	}
}

func getReceiveChannel(ch *amqp.Channel, qName string) (messageChannel <-chan amqp.Delivery) {
	messageChannel, err := ch.Consume(
		qName, // q name
		"",    // consumer
		true,  // autoAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // args
	)
	failOnError(err, "error consuming")
	return
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
