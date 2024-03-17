package routing

import (
	"github.com/arvin-gao/gotutorial/packages/rabbitmq"
	. "github.com/arvin-gao/gotutorial/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type queueBind struct {
	name       string
	bindingKey string // Topic
}

func initExchange(exchangeName string, queues ...queueBind) *amqp.Connection {
	conn := rabbitmq.ConnMQ()

	ch, err := conn.Channel()
	CheckError(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeTopic, // exchange type to `topic`
		true,
		false,
		false,
		false,
		nil,
	)
	CheckError(err)

	for _, queue := range queues {
		_, err := ch.QueueDeclare(
			queue.name,
			false,
			true,
			false,
			false,
			nil,
		)
		CheckError(err)

		err = ch.QueueBind(
			queue.name,
			queue.bindingKey,
			exchangeName,
			false,
			nil,
		)
		CheckError(err)
	}

	return conn
}
