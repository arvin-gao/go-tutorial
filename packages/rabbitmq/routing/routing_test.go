// Situation: We will be able to direct only critical error messages
// to the log file (to save disk space), while still being able to
// print all of the log messages on the console.

package routing

import (
	"context"
	"fmt"
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

	exchangeName := "logs_direct"
	// initialize logs exchange.
	queues := []queueBind{
		{
			"q1",
			"bk1",
		},
		{
			"q2",
			"bk1",
		},
		{
			"q3",
			"bk2",
		},
	}
	fmt.Println("q1 and q2 binding on bk1, but q3 is bk2")
	initConn := initExchange(exchangeName, queues...)
	defer initConn.Close()
	fmt.Println("init exchange complete")

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
		exchangeName,        // name
		amqp.ExchangeDirect, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	CheckError(err)
	for i := 0; i < 10; i++ {
		emitLog(ctx, ch, exchangeName, queues[i%n].bindingKey, rabbitmq.LogInfo(i))
	}
	fmt.Println("emit logs complete.")

	<-ctx.Done()
}
