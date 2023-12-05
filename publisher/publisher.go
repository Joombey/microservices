package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	connection "farukh.go/micro/connection"
	"farukh.go/micro/consts"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, ch := connection.DeclareConnectionAndCreateChannel()
	defer ch.Close()
	defer conn.Close()
	body := "Hello World!"
	for i := 0; i < 5; i++ {
		publish(ch, consts.Q_NAME, []byte(body), rand.Int63n(10))
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func publish(ch *amqp.Channel, qName string, body []byte, timeToSleep int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		consts.EXCHANGE_NAME,                     // exchange
		fmt.Sprintf("exclusive/%d", timeToSleep), // routing key
		false,                                    // mandatory
		false,                                    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
