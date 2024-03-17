package rabbitmq

import (
	"context"
	"log"
	"strconv"
	"time"

	. "github.com/arvin-gao/gotutorial/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection

const queueName = "simple-queue"

func simpleMQ() {
	conn = ConnMQ()
	defer conn.Close()

	ch, err := conn.Channel()
	CheckErrorWithMessage(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	CheckErrorWithMessage(err, "Failed to declare a queue")

	sendMessage(ch, q.Name, LogInfo(1))
	sendMessage(ch, q.Name, LogInfo(2))

	time.Sleep(2 * time.Second)

	receiveMessage(q.Name)
}

func sendMessage(ch *amqp.Channel, queue, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	CheckErrorWithMessage(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", msg)
}

func receiveMessage(queue string) {
	ch, err := conn.Channel()
	CheckErrorWithMessage(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	CheckErrorWithMessage(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func LogInfo(i int) string {
	return "log-info(" + strconv.FormatInt(int64(i), 10) + ")"
}
