package publishsubscribe

import (
	"context"

	. "github.com/arvin-gao/gotutorial/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func emitLog(ctx context.Context, ch *amqp.Channel, exchangeName, msg string) {
	err := ch.PublishWithContext(ctx,
		exchangeName,
		"", // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	CheckError(err)
}
