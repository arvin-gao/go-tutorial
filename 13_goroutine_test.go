package gotutorial

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

var (
	unbuffCh = make(chan int)
	buffCh   = make(chan int, 1)
)

func TestGoroutine(t *testing.T) {
	f := func(str string) {
		time.Sleep(50 * time.Millisecond)
		for i := 0; i < 3; i++ {
			ptr(str, i)
		}
	}

	go f("f1")
	go f("f2")
	go f("f3")
	go f("f4")
	ptr("sleeping...")
	time.Sleep(time.Second)
	ptr("done")
}

func TestGoroutineWithWaitGroup(t *testing.T) {
	f := func(wg *sync.WaitGroup, str string) {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		for i := 0; i < 3; i++ {
			ptr(str, i)
		}
	}

	var wg sync.WaitGroup
	wg.Add(4)
	go f(&wg, "f1")
	go f(&wg, "f2")
	go f(&wg, "f3")
	go f(&wg, "f4")
	wg.Wait()
	ptr("done")
}

func TestCounterWithMultiGoroutine(t *testing.T) {
	var count int

	f := func(wg *sync.WaitGroup, lock *sync.Mutex) {
		defer wg.Done()

		lock.Lock()
		defer lock.Unlock()

		count++
	}

	var wg sync.WaitGroup
	loopCount := 100000

	wg.Add(loopCount)

	var lock sync.Mutex
	for i := 0; i < loopCount; i++ {
		go f(&wg, &lock)
	}

	wg.Wait()

	assert.Equal(t, loopCount, count)
}

func TestCounterWithGoroutineAndLock(t *testing.T) {
	var (
		lock  sync.Mutex
		count int
	)

	f := func(wg *sync.WaitGroup, lock *sync.Mutex) {
		defer wg.Done()

		lock.Lock()
		defer lock.Unlock()

		count++
	}

	var wg sync.WaitGroup
	loopCount := 100000

	wg.Add(loopCount)

	for i := 0; i < loopCount; i++ {
		go f(&wg, &lock)
	}

	wg.Wait()

	assert.Equal(t, loopCount, count)
}

func TestGoroutineExample(t *testing.T) {
	var v1 = 123
	go func(v2 int) {
		time.Sleep(time.Second)
		ptr(v1, v2)
	}(v1)

	v1 = 789

	time.Sleep(2 * time.Second)
}

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

/*
	Unbuffered will only accept sends (chan <-) if there is a corresponding receive
	(<- chan) ready to receive the sent value.
	Buffered channels accept a limited number of values without a corresponding
	receiver for those values.
*/

/*
unbuffered channel
	sender 會等待(ch <- value)
	receiver 會等待(v, isOpen := <-ch)
buffered channel
	sender 會無視等待，除非該 channel buffer 已滿
	receiver 會等待
*/

func TestChannelWithDeadlock(t *testing.T) {
	// fatal error: all goroutines are asleep - deadlock!
	unbuffCh <- 1
	ptr(<-unbuffCh)
}

func TestChannelWithDeadlock2(t *testing.T) {
	buffCh <- 1
	// fatal error: all goroutines are asleep - deadlock!
	buffCh <- 1
	ptr(<-buffCh)
}

// if send a msg to closed channel, that will be panic with 'panic: send on closed channel' message.
func TestChannelWithClose(t *testing.T) {
	buffCh <- 1
	ptr(unsafe.Sizeof(buffCh))
	close(buffCh)
	// buffCh <- 1
	ptr(unsafe.Sizeof(buffCh))

}

func TestChannelWithLoop(t *testing.T) {
	for c := range unbuffCh {
		ptr("1-for:", c)
	}

	// check the channel status.
	for {
		if c, open := <-unbuffCh; !open {
			ptr("2-for:", c)
		}
	}
}

// for range 會持續讀資料直到 channel 被關閉為止
func TestChannelWithRangeLoop(t *testing.T) {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		ptr(elem)
	}
}

func TestSingleChannelWithGoroutineLoop(t *testing.T) {
	count := 10
	go func() {
		for i := 0; i < count; i++ {
			go func(n int) {
				buffCh <- n
				ptr("bufferedCh <- n:", n)
			}(i)
		}
	}()

	var i int
	for v := range buffCh {
		ptr("v:", v)
		time.Sleep(500 * time.Millisecond)
		i++
		if i == count {
			close(buffCh)
		}
	}

	// check the channel status.
	if v, open := <-buffCh; open {
		ptr("is opening. v:", v)
	} else {
		ptr("not open")
	}
}

// This specificity increases the type-safety of the program
func TestChannelWithChannelParamFunc(t *testing.T) {
	// Only accepts a channel for sending values.
	ping := func(pings chan<- string, msg string) {
		pings <- msg
		// <-pings // invalid code
	}

	// Accepts one channel for receives (pings) and a second for sends (pongs)
	pong := func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
		// <-pongs // invalid code
	}

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go ping(pings, "passed message")
	go pong(pings, pongs)

	ptr(<-pongs)
}

/*
Go’s select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.
*/
func TestChannelWithSelect(t *testing.T) {
	ptr("setup unbuffered channel")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	pTitle("select with loop")
	var isBreak bool
	for {
		if isBreak {
			break
		}

		select {
		case msg1 := <-c1:
			ptr("received", msg1)
		case msg2 := <-c2:
			ptr("received", msg2)
		case <-time.After(2 * time.Second):
			ptr("timeout")
			isBreak = true
		}
	}

	ptr()
	ptr("setup buffered channel")
	c1 = make(chan string, 1)
	c2 = make(chan string, 1)

	c1 <- "one"
	c2 <- "two"

	pTitle("single select")
	select {
	case msg1 := <-c1:
		ptr("received", msg1)
	case msg2 := <-c2:
		ptr("received", msg2)
	}

	pTitle("select with loop")
	c1 = make(chan string, 1)
	c2 = make(chan string, 1)
	c1 <- "one"
	c2 <- "two"
	for {
		select {
		case msg1 := <-c1:
			ptr("received", msg1)
		case msg2 := <-c2:
			ptr("received", msg2)
		case <-time.After(100 * time.Millisecond):
			return
		}
	}
}

type Counter struct {
	sync.Mutex
	v int
}

func (c *Counter) CountWithoutLock() {
	c.v++
}

func (c *Counter) CountWithLock() {
	c.Lock()
	defer c.Unlock()
	c.v++
}

func TestSyncWithWaitGroupAndMutex(t *testing.T) {
	var wg sync.WaitGroup
	var c1, c2 Counter

	count := 1000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			c1.CountWithoutLock()
			c2.CountWithLock()
		}(&wg)
	}
	wg.Wait()

	ptr("counter-1:", c1.v)
	ptr("counter-2(lock):", c2.v)
}

func TestChannelExample(t *testing.T) {
	jobs := make(chan int, 1)
	done := make(chan bool)

	go func() {
		for {
			v, isOpen := <-jobs
			if isOpen {
				ptr("received job", v)
			} else {
				ptr("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		ptr("sent job", j)
	}

	close(jobs)
	ptr("sent all jobs")

	<-done
}

// Timers are for when you want to do something
// once in the future.
func TestChannelWithTimer(t *testing.T) {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	ptr("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		ptr("Timer 2 fired")
	}()

	stop := timer2.Stop()
	if stop {
		ptr("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

// Tickers are for when you want to do something
// repeatedly at regular intervals.
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				ptr("done")
				return
			case t := <-ticker.C:
				ptr("Tick at", t.Format(time.DateTime))
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	ticker.Stop()
	done <- true
	ptr("Ticker stopped")
}

func TestChannelExample2(t *testing.T) {
	var sum int

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	select {
	case v := <-ch1:
		sum += v
	case v := <-ch2:
		sum += v
	case <-time.After(100 * time.Millisecond):
		ptr("count:", sum)
		return
	}

	ptr("count:", sum)
}

func TestChannelExample3(t *testing.T) {
	sum := 0
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	for {
		select {
		case v := <-ch1:
			sum += v
		case v := <-ch2:
			sum += v
		case <-time.After(100 * time.Millisecond):
			ptr("count:", sum)
			return
		}
	}
}

func TestWorkerWithGoroutineAndChannel(t *testing.T) {
	worker := func(id int, jobs <-chan int, results chan<- int) {
		for j := range jobs {
			pf("worker-%d started  job-%d", id, j)
			time.Sleep(time.Second)
			pf("worker-%d finished  job-%d", id, j)
			results <- j * 2
		}
	}

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		ptr("got result:", <-results)
	}
}

// example from https://blog.wu-boy.com/2022/05/read-data-from-channel-in-go/
func TestGoroutineExample1(t *testing.T) {
	str := []byte("foobar")

	ch := make(chan byte, len(str))
	next := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < len(str); i++ {
		ch <- str[i]
	}

	close(ch)

	go func() {
		defer wg.Done()
		for {
			if _, ok := <-next; !ok {
				return
			}

			v, ok := <-ch
			if !ok {
				close(next)
				return
			}

			ptr("goroutine01:", string(v))

			next <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if _, ok := <-next; !ok {
				return
			}

			v, ok := <-ch
			if !ok {
				close(next)
				return
			}

			ptr("goroutine02:", string(v))

			next <- struct{}{}
		}
	}()

	next <- struct{}{}

	wg.Wait()
}

// TODO: .
func TestLockIncreaseCountByAtomic(t *testing.T) {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				// ops++ // normal increment.
				atomic.AddUint64(&ops, 1) // lock increment.
			}
		}()
	}

	wg.Wait()

	ptr("ops:", ops)
}
