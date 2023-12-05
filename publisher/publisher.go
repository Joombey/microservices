package main

import (
	"context"
	"log"
	"time"

	connection "farukh.go/micro/connection"
	"farukh.go/micro/consts"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, ch := connection.DeclareConnectionAndCreateChannel()
	defer ch.Close()
	defer conn.Close()
	qName := connection.DeclareQWithBinding(ch)
	body := "Hello World!"
	publish(ch, qName, []byte(body))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func publish(ch *amqp.Channel, qName string, body []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		consts.EXCHANGE_NAME, // exchange
		"",                   // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
