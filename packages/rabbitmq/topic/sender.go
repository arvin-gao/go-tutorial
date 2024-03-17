package routing

import (
	"context"

	. "github.com/arvin-gao/gotutorial/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

func emitLog(ctx context.Context, ch *amqp.Channel, exchangeName, routingKey, msg string) {
	err := ch.PublishWithContext(ctx,
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	CheckError(err)
}
