package workqueues

// Work Queues (aka: Task Queues)
// time-consuming tasks among multiple workers.

// We encapsulate a task as a message and send it
// to a queue. A worker process running in the
// background will pop the tasks and eventually
// execute the job. When you run many workers the
// tasks will be shared between them.

/*
            -> C1
          /
P -> Queue
          \
					 ->  C2
*/

// One of the advantages of using a Task Queue is the
// ability to easily parallelise work.
// If we are building up a backlog of work, we can just
// add more workers and that way, scale easily.

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/arvin-gao/gotutorial/packages/rabbitmq"
	"github.com/arvin-gao/gotutorial/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

const queueName = "task-queue"

var conn *amqp.Connection

func taskQueue() {
	conn = rabbitmq.ConnMQ()
	defer conn.Close()

	sendTask("START TASK WORK")
	fmt.Println("build workers")
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 2; i++ {
		go worker(ctx, i+1)
	}
	for i := 0; i < 1000; i++ {
		sendTask("msg-" + strconv.FormatInt(int64(i), 10))
	}
	fmt.Println("send messages done")
	time.Sleep(5 * time.Second)
	cancel()
}

func sendTask(msg string) {
	ch, err := conn.Channel()
	utils.CheckError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	utils.CheckError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(msg),
		})
	utils.CheckError(err)
}

func worker(ctx context.Context, id int) {
	ch, err := conn.Channel()
	utils.CheckError(err)
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	utils.CheckError(err)

	workerName := "worker-" + strconv.FormatInt(int64(id), 10)
	go func() {
		for d := range msgs {
			log.Printf("%s consumes: %s", workerName, d.Body)
			utils.CheckError(d.Ack(false))
			time.Sleep(time.Duration(rand.Intn(600)) * time.Millisecond)
		}
	}()
	log.Println(workerName, "build success")
	<-ctx.Done()
	ch.Cancel(queueName, true)
	log.Println(workerName, "closed")
}
