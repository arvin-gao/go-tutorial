package gotutorial

import (
	"testing"
	"time"
)

func TestLoop(t *testing.T) {
	var i int

	// Basic loop
	for j := 7; j <= 15; j++ {
		// Support `break` statement.
		if i > 10 {
			break
		}
		// Support `continue` statement.
		if i == 5 {
			continue
		}
	}

	// Infinite loop
	for {
		break
	}
	// Infinite loop.
	// Similar to `for i := 0; true; i++`
	for i := 0; ; i++ {
		break
	}

	// While loop.
	i = 0
	for i <= 3 {
		i = i + 1
	}
	// While loop.
	i = 0
	for {
		if i <= 3 {
			break
		}
		i++
	}

	// Loop with range slice or array.
	nums := []int{1, 2, 3}
	// Get the index and value while looping through the range
	for i, num := range nums {
		ptrlnf("index:%d value:%d", i, num)
	}

	// Only getting the index.
	for i := range nums {
		ptr(i)
	}

	// Only getting the value.
	for _, num := range nums {
		ptr(num)
	}
}

func TestLoopMap(t *testing.T) {
	m := map[string]string{
		"1": "v1",
		"2": "v2",
		"3": "v3",
		"4": "v4",
	}
	// 順序不一定
	for key, value := range m {
		ptr(key, value)
	}

	for key := range m {
		ptr(key)
	}
}

func TestLoopChannel(t *testing.T) {
	ch := make(chan string)

	// Close the channel to ensure process done.
	go func() {
		<-time.After(time.Second)
		close(ch)
	}()

	ch <- "data"
	for x := range ch {
		ptr(x)
	}
}
