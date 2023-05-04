package gotutorial

import (
	"fmt"
	"sync"
	"testing"
)

type Counter struct {
	sync.RWMutex
	v int
}

func (c *Counter) Count() {
	c.v++
}

func (c *Counter) CountWithLock() {
	c.Lock()
	defer c.Unlock()
	c.v++
}

func TestSync(t *testing.T) {
	var wg sync.WaitGroup
	var c1, c2 Counter

	count := 1000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			c1.Count()
			c2.CountWithLock()
		}(&wg)
	}
	wg.Wait()

	fmt.Println("counter-1:", c1.v)
	fmt.Println("counter-2(lock):", c2.v)
}
