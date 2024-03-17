package routing

import (
	"github.com/arvin-gao/gotutorial/packages/rabbitmq"
	. "github.com/arvin-gao/gotutorial/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type queueBind struct {
	name       string
	bindingKey string
}

func initExchange(exchangeName string, queues ...queueBind) *amqp.Connection {
	conn := rabbitmq.ConnMQ()

	ch, err := conn.Channel()
	CheckError(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeDirect, // exchange type to `direct`
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
			queue.name,       // queue name
			queue.bindingKey, // binding key as routing key.
			exchangeName,     // exchange name
			false,
			nil,
		)
		CheckError(err)
	}

	return conn
}
