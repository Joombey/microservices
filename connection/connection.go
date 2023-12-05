package connection

import (
	"log"

	"farukh.go/micro/consts"
	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareConnectionAndCreateChannel() (conn *amqp.Connection, ch *amqp.Channel) {
	conn, err := amqp.Dial(consts.AMQP_URL)
	defer failOnError(err, "failed to decalre connection")
	ch, err = conn.Channel()
	defer failOnError(err, "failed to create channel")
	return
}

func DeclareQ(ch *amqp.Channel) (q amqp.Queue) {
	q, err := ch.QueueDeclare(
		consts.Q_NAME, // q name
		true,          // durable
		false,         // autoDetele
		false,         // exclusive
		false,         // noWait
		nil,           // args
	)
	failOnError(err, "failed to create Queue")
	return
}

func DelcareExcange(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		consts.EXCHANGE_NAME, //name
		amqp.ExchangeFanout,  //kind
		true,                 //durable
		false,                //autoDelete
		false,                //internal
		false,                //noWait
		nil,                  //args
	)
	failOnError(err, "failed to create exchange")
}

func DeclareQWithBinding(ch *amqp.Channel) (q string) {
	DelcareExcange(ch)
	q = DeclareQ(ch).Name
	err := ch.QueueBind(
		q,                    // q_name
		consts.ROUTING_KEY,   // key - idk
		consts.EXCHANGE_NAME, // exchange_name
		false,                // noWait
		nil,                  // args
	)
	defer failOnError(err, "failed to bind q")
	return
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
