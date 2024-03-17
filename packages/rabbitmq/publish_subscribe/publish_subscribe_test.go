package publishsubscribe

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

	exchangeName := "logs"
	// initialize logs exchange.
	initConn, queue := initExchange(exchangeName)
	defer initConn.Close()
	fmt.Println("init exchange complete")

	// run three receiver.
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func(id int) {
			consumer(ctx, &wg, id, queue)
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
		amqp.ExchangeFanout, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	CheckError(err)
	for i := 0; i < 10; i++ {
		emitLog(ctx, ch, exchangeName, rabbitmq.LogInfo(i))
	}
	fmt.Println("emit logs complete.")

	<-ctx.Done()
}
