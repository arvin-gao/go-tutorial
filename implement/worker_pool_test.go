package implement

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestWorkerPool(t *testing.T) {
	WorkerPool()
}

type TaskFunc func(ctx context.Context) error

func WorkerPool() {
	tasks := make([]TaskFunc, 100)
	for i := 0; i < len(tasks); i++ {
		func(index int) {
			tasks[index] = func(ctx context.Context) error {
				fmt.Println(index)
				if index == 51 {
					return errors.New("missing")
				}
				return nil
			}
		}(i)
	}

	if err := ExecuteAll(0, tasks...); err != nil {
		logrus.Error(err)
	}
}

// numCPU is runtime.NumCPU() if the numCPU is zero.
func ExecuteAll(numCPU int, tasks ...TaskFunc) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error

	if numCPU == 0 {
		numCPU = runtime.NumCPU()
	}

	var wg sync.WaitGroup
	wg.Add(numCPU)

	queue := make(chan TaskFunc, len(tasks))

	// Spawn the executer
	for i := 0; i < numCPU; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case task, ok := <-queue:
					if ctx.Err() != nil || !ok {
						return
					}
					println("get task")
					if taskErr := task(ctx); taskErr != nil {
						err = taskErr
						cancel()
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	for _, task := range tasks {
		queue <- task
	}
	close(queue)

	wg.Wait()
	return err
}
