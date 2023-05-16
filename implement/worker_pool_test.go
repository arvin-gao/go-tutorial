package implement

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	WorkerPool()
}

type TaskFunc func(ctx context.Context) error

func WorkerPool() {
	tasks := make([]TaskFunc, 0, 100)
	for i := 0; i < 100; i++ {
		func(val int) {
			tasks = append(tasks, func(ctx context.Context) error {
				fmt.Println(val)
				if val == 51 {
					return errors.New("missing")
				}
				return nil
			})
		}(i)
	}

	err := ExecuteAll(0, tasks...)
	if err == nil {
		fmt.Println("missing error")
	}
}

func ExecuteAll(numCPU int, tasks ...TaskFunc) error {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	if numCPU == 0 {
		numCPU = runtime.NumCPU()
	}
	fmt.Println("numCPU:", numCPU)
	queue := make(chan TaskFunc, numCPU)

	// Spawn the executer
	for i := 0; i < numCPU; i++ {
		go func() {
			for task := range queue {
				fmt.Println("get task")
				if err == nil {
					taskErr := task(ctx)
					if taskErr != nil {
						err = taskErr
						cancel()
					}
				}
				wg.Done()
			}
		}()
	}

	// Add tasks to queue
	for _, task := range tasks {
		queue <- task
	}
	close(queue)

	// wait for all task done
	wg.Wait()
	return err
}
