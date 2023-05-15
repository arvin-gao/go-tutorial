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
	taskCount := 100

	tasks := make([]TaskFunc, taskCount, taskCount)

	for i := 0; i < taskCount; i++ {
		func(val int) {
			tasks[i] = func(ctx context.Context) error {
				fmt.Println(i)
				if i == 51 {
					return errors.New("missing")
				}
				return nil
			}
		}(i)
	}

	if err := ExecuteAll(0, tasks...); err == nil {
		fmt.Println(err)
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
					if taskErr := task(ctx); taskErr != nil {
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
