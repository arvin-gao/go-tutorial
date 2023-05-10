package gotutorial

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {
	pTitle("limiter")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		println("request", req, time.Now())
	}

	pTitle("burstyLimiter")
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		println("request", req, time.Now())
	}
}

// TODO: learning.
func TestIncreaseCounterByAtomic(t *testing.T) {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				// ops++
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	wg.Wait()

	println("ops:", ops)
}
