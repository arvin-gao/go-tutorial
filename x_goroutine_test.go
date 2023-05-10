package gotutorial

import (
	"sync"
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
		for i := 0; i < 3; i++ {
			println(str, ":", i)
		}
	}

	f("f1")

	go f("f2")

	go func(msg string) {
		println(msg)
	}("f3")

	time.Sleep(time.Second)
	println("done")
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
	println(<-unbuffCh)
}

func TestChannelWithDeadlock2(t *testing.T) {
	buffCh <- 1
	// fatal error: all goroutines are asleep - deadlock!
	buffCh <- 1
	println(<-buffCh)
}

// if send a msg to closed channel, that will be panic with 'panic: send on closed channel' message.
func TestChannelWithClose(t *testing.T) {
	buffCh <- 1
	println(unsafe.Sizeof(buffCh))
	close(buffCh)
	// buffCh <- 1
	println(unsafe.Sizeof(buffCh))

}

func TestChannelWithLoop(t *testing.T) {
	for c := range unbuffCh {
		println("1-for:", c)
	}

	// check the channel status.
	for {
		if c, open := <-unbuffCh; !open {
			println("2-for:", c)
		}
	}
}

func TestChannelWithRangeLoop(t *testing.T) {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		println(elem)
	}
}

func TestSingleChannelWithGoroutineLoop(t *testing.T) {
	count := 10
	go func() {
		for i := 0; i < count; i++ {
			go func(n int) {
				buffCh <- n
				println("bufferedCh <- n:", n)
			}(i)
		}
	}()

	var i int
	for v := range buffCh {
		println("v:", v)
		time.Sleep(500 * time.Millisecond)
		i++
		if i == count {
			close(buffCh)
		}
	}

	// check the channel status.
	if v, open := <-buffCh; open {
		println("is opening. v:", v)
	} else {
		println("not open")
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

	println(<-pongs)
}

/*
Go’s select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.
*/
func TestChannelWithSelect(t *testing.T) {
	println("setup unbuffered channel")
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
			println("received", msg1)
		case msg2 := <-c2:
			println("received", msg2)
		case <-time.After(2 * time.Second):
			println("timeout")
			isBreak = true
		}
	}

	println()
	println("setup buffered channel")
	c1 = make(chan string, 1)
	c2 = make(chan string, 1)

	c1 <- "one"
	c2 <- "two"

	pTitle("single select")
	select {
	case msg1 := <-c1:
		println("received", msg1)
	case msg2 := <-c2:
		println("received", msg2)
	}

	pTitle("select with loop")
	c1 = make(chan string, 1)
	c2 = make(chan string, 1)
	c1 <- "one"
	c2 <- "two"
	for {
		select {
		case msg1 := <-c1:
			println("received", msg1)
		case msg2 := <-c2:
			println("received", msg2)
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

	println("counter-1:", c1.v)
	println("counter-2(lock):", c2.v)
}

func TestChannelExample(t *testing.T) {
	jobs := make(chan int, 1)
	done := make(chan bool)

	go func() {
		for {
			v, isOpen := <-jobs
			if isOpen {
				println("received job", v)
			} else {
				println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		println("sent job", j)
	}

	close(jobs)
	println("sent all jobs")

	<-done
}

// Timers are for when you want to do something
// once in the future.
func TestChannelWithTimer(t *testing.T) {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		println("Timer 2 fired")
	}()

	stop := timer2.Stop()
	if stop {
		println("Timer 2 stopped")
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
				println("done")
				return
			case t := <-ticker.C:
				println("Tick at", t.Format(time.DateTime))
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	ticker.Stop()
	done <- true
	println("Ticker stopped")
}

func TestAA(t *testing.T) {
	var lock sync.Locker

	go func() {
		for i := 0; i < 1000; i++ {
			func() {
				lock.Lock()
				defer lock.Unlock()
				count := 0
				ch1 := make(chan int, 1)
				ch2 := make(chan int, 1)
				for {
					select {
					case <-ch1:
						count++
					case <-ch2:
						count++
					}
				}
			}()
		}

	}()

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
		println("count:", sum)
		return
	}

	println("count:", sum)
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
			println("count:", sum)
			return
		}
	}
}

func TestTestee(t *testing.T) {
	var (
		lock               sync.Mutex
		oneCount, twoCount int
	)

	f := func(wg *sync.WaitGroup, lock *sync.Mutex) {
		lock.Lock()

		defer wg.Done()
		defer lock.Unlock()

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
			println("timeout. count:", sum)
			return
		}

		if sum == 1 {
			oneCount++
		} else if sum == 2 {
			twoCount++
		} else {
			println("invalid sum:", sum)
		}
	}

	var wg sync.WaitGroup
	loopCount := 100000

	wg.Add(loopCount)

	for i := 0; i < loopCount; i++ {
		go f(&wg, &lock)
	}

	wg.Wait()

	pf("one count: %d\ntwo count: %d", oneCount, twoCount)

	assert.Equal(t, loopCount, oneCount+twoCount)
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
		println("got result:", <-results)
	}
}
