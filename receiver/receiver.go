package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

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

		valuePart := strings.Split(delivery.RoutingKey, "/")[1]
		timeToSleep, _ := strconv.ParseInt(valuePart, 10, 64)

		fmt.Printf("i need to slepp for %d\n", timeToSleep)

		time.Sleep(time.Duration(timeToSleep) * time.Second)

		delivery.Ack(false)

		fmt.Printf("%d seconds elapsed\n", timeToSleep)
	}
}

func getReceiveChannel(ch *amqp.Channel, qName string) (messageChannel <-chan amqp.Delivery) {
	messageChannel, err := ch.Consume(
		qName, // q name
		"",    // consumer
		false, // autoAck
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
