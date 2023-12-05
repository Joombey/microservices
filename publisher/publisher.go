package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	connection "farukh.go/micro/connection"
	"farukh.go/micro/consts"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, ch := connection.DeclareConnectionAndCreateChannel()
	defer ch.Close()
	defer conn.Close()
	connection.DeclareQWithBinding(ch)
	for i := 0; i < 5; i++ {
		randomNumber := rand.Int31n(10)
		body := strconv.Itoa(int(randomNumber))
		publish(ch, consts.Q_NAME, []byte(body), randomNumber)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func publish(ch *amqp.Channel, qName string, body []byte, timeToSleep int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		consts.EXCHANGE_NAME,                     // exchange
		fmt.Sprintf("exclusive/%d", timeToSleep), // routing key
		false,                                    // mandatory
		false,                                    // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // delivery mode
			ContentType:  "text/plain",    // content type
			Body:         body,            // content
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
