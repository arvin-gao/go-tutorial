package publishsubscribe

import (
	"github.com/arvin-gao/gotutorial/packages/rabbitmq"
	. "github.com/arvin-gao/gotutorial/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func initExchange(exchangeName string) (conn *amqp.Connection, queue string) {
	conn = rabbitmq.ConnMQ()

	ch, err := conn.Channel()
	CheckError(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,        // exchange name
		amqp.ExchangeFanout, // exchange type
		true,
		false,
		false,
		false,
		nil,
	)
	CheckError(err)

	q, err := ch.QueueDeclare(
		"", // empty queue name
		false,
		true,
		false,
		false,
		nil,
	)
	CheckError(err)

	err = ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange name
		false,
		nil,
	)
	CheckError(err)

	queue = q.Name
	return
}
