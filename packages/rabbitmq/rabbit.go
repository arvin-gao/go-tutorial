package rabbitmq

import (
	. "github.com/arvin-gao/gotutorial/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

const dialURL = "amqp://root:1234@localhost:5672/"

func ConnMQ() *amqp.Connection {
	conn, err := amqp.Dial(dialURL)
	CheckErrorWithMessage(err, "connect failed")
	return conn
}
