package routing

import (
	"context"
	"fmt"
	"sync"

	. "github.com/arvin-gao/gotutorial/utils"

	"github.com/arvin-gao/gotutorial/packages/rabbitmq"
)

func consumer(ctx context.Context, wg *sync.WaitGroup, queueIdx int, queue string) {
	conn := rabbitmq.ConnMQ()
	defer conn.Close()

	ch, err := conn.Channel()
	CheckError(err)
	defer ch.Close()

	msgs, err := ch.Consume(
		queue, // queue name(random-name)
		"",
		true, // auto-ack
		false,
		false,
		false,
		nil,
	)
	CheckError(err)

	go func() {
		wg.Done()
		for msg := range msgs {
			fmt.Printf("Queue(%d); Content: %s\n", queueIdx, msg.Body)
		}
	}()
	<-ctx.Done()
}
