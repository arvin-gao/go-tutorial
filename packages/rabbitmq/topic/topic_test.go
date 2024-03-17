// Situation: We will be able to direct only critical error messages
// to the log file (to save disk space), while still being able to
// print all of the log messages on the console.

package routing

import (
	"context"
	"sync"
	"testing"
	"time"

	. "github.com/arvin-gao/gotutorial/utils"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/arvin-gao/gotutorial/packages/rabbitmq"
)

var (
	timeout = 5 * time.Second
)

func TestEmitLogs(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	exchangeName := "logs_topic"
	// initialize logs exchange.
	queues := []queueBind{
		{
			"q0",
			"house.bedroom.*.#",
		},
		{
			"q1",
			"house.room.device.*",
		},
		{
			"q2",
			"*.room.*.furniture.#",
		},
	}
	initConn := initExchange(exchangeName, queues...)
	defer initConn.Close()

	n := len(queues)
	// run three receiver.
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			consumer(ctx, &wg, idx, queues[idx].name)
		}(i)
	}
	wg.Wait()

	// send fake logging informations.
	conn := rabbitmq.ConnMQ()
	defer conn.Close()
	ch, err := conn.Channel()
	CheckError(err)
	defer ch.Close()
	err = ch.ExchangeDeclare(
		exchangeName,       // name
		amqp.ExchangeTopic, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	CheckError(err)

	messages := []struct {
		topic   string
		content string
	}{
		{"house", "The `house` doesn't match any queues"},
		{"house.bathroom.toilet", "The `house.bathroom.toilet` doesn't match any queues"},
		{"house.bedroom", "The `house.bathroom` doesn't match any queues"},
		{"house.bedroom.bed", "The `house.bedroom.bed` will match `Q0`"},
		{"house.room.device.pc", "The `house.room.device.pc` will match `Q1`"},
		{"house.room.device.phone", "The `house.room.device.phone` will match `Q1`"},
		{"house.room.device.furniture", "The `house.room.device.furniture` will match `Q1` and `Q2`"},
		{"house2.room.r1.furniture.bed", "The `house2.room.r1.furniture.bed` will match `Q2`"},
		{"house3.room.r2.furniture.desk", "The `house3.room.r2.furniture.desk` will match `Q2`"},
	}
	for _, msg := range messages {
		go emitLog(ctx, ch, exchangeName, msg.topic, msg.content)
	}

	<-ctx.Done()
}
