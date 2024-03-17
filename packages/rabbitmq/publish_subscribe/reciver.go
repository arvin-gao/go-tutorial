package publishsubscribe

import (
	"context"
	"fmt"
	"sync"
	"time"

	. "github.com/arvin-gao/gotutorial/utils"

	"github.com/arvin-gao/gotutorial/packages/rabbitmq"
)

func consumer(ctx context.Context, wg *sync.WaitGroup, id int, queue string) {
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
		for msg := range msgs {
			fmt.Printf("ID(%d) Received: %s\n", id, msg.Body)
		}
	}()
	fmt.Println("consumer is running")
	wg.Done()
	<-ctx.Done()
	return

	go func() {
		for {
			select {
			case msg, ok := <-msgs:
				if ok {
					fmt.Printf("ID:%d, Received: %s\n", id, msg.Body)
				}
				// msg.Ack(false)
			case <-ctx.Done():
				fmt.Printf("Consumer(%d) done\n", id)
				return
			default:
				fmt.Printf("Consumer(%d) waiting...\n", id)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	fmt.Printf("Consumer(%d) start\n", id)
	wg.Done()
}
