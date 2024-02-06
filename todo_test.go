package gotutorial

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// ?rename
// defer -> 10_function -> second layer
// goroutine-> second layer -> channel.go

// ? struct
// 封裝(Encapsulation)
/*
通過限制只有特定類別的物件可以存取這一特定類別的成員
存取權限：public, private
*/

// 繼承(extends, override)

// 多型(Polymorphism
/*
abstract, interface
*/

// TODO: tidy.
/*
1. Field C.A1.g and method C.B.g collide, so they are both not promoted.
2. Method C.B.f gets promoted as C.f.
3. Method C.m overrides C.A1.m.
*/

type A1 struct {
	g int
}

func (A1) m() int {
	return 1
}

type B int

func (B) g() {
	// pass
}

func (B) f() {
	// pass
}

type C struct {
	A1
	B
}

func (C) m() int {
	return 9
}

func TestEmbed(t *testing.T) {
	var c interface{} = C{}
	//  Method C.B.f gets promoted as C.f.
	_, bf := c.(interface{ f() })
	// Field C.A1.g and method C.B.g collide, so they are both not promoted.
	_, bg := c.(interface{ g() })
	// Method C.m overrides C.A1.m.
	i := c.(interface{ m() int })
	ptr(bf, bg, i.m())
}

//? regex
// https://willh.gitbook.io/build-web-application-with-golang-zhtw/07.0/07.3

// ?file.go
// TODO
func OtherDirectories(path string) {
	// Changes the current working directory to the named directory.
	os.Chdir(path)

	// ?
	visit := func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ptr(" ", p, info.IsDir())
		return nil
	}
	filepath.Walk("subdir", visit)
}

// ?error.go

// TODO: more info(errors.Is and errors.As).
// TODO: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

// ?goroutine.go
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

// TODO
func TestSearchExample(t *testing.T) {
	result := make(chan int, 4)
	replicas := make([]replica, 4)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < len(replicas); i++ {
		go func(i int) {
			result <- replicas[i].searchData(ctx, i)
		}(i)
	}
	fmt.Println("search result:", <-result)
	cancel()
	time.Sleep(1 * time.Second)
}

type replica struct{}

func (r replica) searchData(ctx context.Context, i int) int {
	f := func(i int) <-chan int {
		c := make(chan int, 1)
		time.Sleep(time.Duration(100*(i+1)) * time.Millisecond)
		c <- 1
		return c
	}
	select {
	case <-f(i):
		fmt.Println("search done:", i)
		return i
	case <-ctx.Done():
		fmt.Println("exit:", i)
		return -1
	}
}

type Result struct{}
type ReplicaSearch func(string) Result

func searchByReplicas(query string, replicas ...ReplicaSearch) Result {
	c := make(chan Result)
	done := make(chan struct{})
	// defer close(done)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
			ptr("search result: ", i)
		case <-done:
			ptr("done:", i)
		case <-time.Tick(100 * time.Millisecond):
			ptr("searching...")
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// The result channel in the searchByReplicas() function is unBuffered.
// This means that only the first goroutine returns.
// All other goroutines are stuck trying to send their results.
// This means if you have more than one replica each call will leak resources.
// Solutions:
//  1. Use a buffered result channel big enough to hold all results.
//  2. Use a select statement with a default case and a buffered result channel
//     that can hold one value.
//     The default case ensures that the goroutines don't get stuck
//     even when the result channel can't receive messages.
func TestSearchByReplicas(t *testing.T) {
	_ = func(query string, replicas ...ReplicaSearch) Result {
		// unBuffered!
		c := make(chan Result)
		searchReplica := func(i int) {
			c <- replicas[i](query)
		}
		for i := range replicas {
			go searchReplica(i)
		}
		// That will stuck all other goroutines.
		return <-c
	}

	//  Solution 1: Use a buffered result channel big enough to hold all results.
	_ = func(query string, replicas ...ReplicaSearch) Result {
		// Using the buffered result channel.
		c := make(chan Result, len(replicas))
		searchReplica := func(i int) {
			c <- replicas[i](query)
		}
		for i := range replicas {
			go searchReplica(i)
		}
		return <-c
	}

	//  Solution 2: Use a select statement with a default case and a buffered result channel
	//  that can hold one value.
	_ = func(query string, replicas ...ReplicaSearch) Result {
		c := make(chan Result, 1)
		searchReplica := func(i int) {
			select {
			case c <- replicas[i](query):
			default:
			}
		}
		for i := range replicas {
			go searchReplica(i)
		}
		return <-c
	}

	// Solution 3: use a special cancellation channel to interrupt the workers.
	_ = func(query string, replicas ...ReplicaSearch) Result {
		c := make(chan Result)
		done := make(chan struct{})
		defer close(done)
		searchReplica := func(i int) {
			select {
			case c <- replicas[i](query):
			case <-done:
			}
		}
		for i := range replicas {
			go searchReplica(i)
		}
		return <-c
	}
}

type Worker struct {
	stop chan struct{}
	done chan struct{}
}

func NewWorker() *Worker {
	w := &Worker{
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}
	return w
}

func (w *Worker) doWork() {
	defer close(w.done)
	for {
		select {
		case <-w.stop:
			println("stop")
			return
		case <-time.After(time.Second):
			println("work!")
		}
	}
}

// Shutdown tells the worker to stop and waits until it has finished.
func (w *Worker) Shutdown() {
	close(w.stop)
	<-w.done
	println("shutdown completely")
}
func Test(t *testing.T) {
	done := make(chan int)
	c := make(chan int)
	go func() {
		count := 0
		for {
			select {
			case <-time.After(500 * time.Millisecond):
				c <- count
				count++
			case <-done:
				close(c)
			}
		}
	}()
	go func() {
		<-time.After(2 * time.Second)
		println("sending done signal!")
		close(done)
	}()

	for v := range c {
		println("rec:", v)
	}
	println("done")
}

// ? variable_const_iota.go
// TODO: tidy
/*
Go spec says: Within a parenthesized const declaration list
the expression list may be omitted from any but the first ConstSpec.

Such an empty list is equivalent to the textual substitution of
the first preceding non-empty expression list and its type if any.
*/

// * iota
/*
The value of the predeclared iota is the constant specification order id (0-based) in a constant declaration.
*/
func TestConstantWithIota(t *testing.T) {
	const (
		_    = 6
		A, _ = iota, iota + 10 // 1, 1 + 10
		_, _                   // 2, 2 + 10
		_, B                   // 3, 3 + 10
	)
	ptr(A, B)
}

// * const mechanism
/*
A local identifier will shadow the global identifier with the same name.

! Notice: The output result was 6 6 when using Go toolchain v1.17-. The bug has been fixed since Go toochain v1.18.
*/

const X = 3

func TestConstantWithGlo(t *testing.T) {
	const (
		X = X + X // X + X. The two "X" are both the `global` one
		Y         // X + X. The two "X" are both the `local` one
	)

	ptr(X, Y)
}

// * ???

/*
Untyped X is a rune constant (in other words, its default type is rune, a.k.a int32, a 32-bit integer type).
Untyped Y is an int constant (in other words, its default type is int, a 64-bit integer type on 64-bit OSes).

At run time, the expression A << n overflows, so it is evaluated as 0; on the other hand, the expression B << n doesn't overflow.
*/
func TestConstantWithOperands(t *testing.T) {
	const (
		X       = '\x61' // 'a'
		Y       = 0x62
		A       = Y - X // 1
		B int64 = 1
	)

	var n = 32
	if A == B {
		ptr(A<<n>>n, B<<n>>n)
	}
}

// variable_map_test.go
// TODO
/*
When using the make function to create a map,
the second argument is neither the initial
length nor the capacity of the result map.

It is just a hint for Go runtime to allocate
a large enough backing array to hold at least
the specified number of entries.

The length (number of entries) of the result
`map` is zero.
m["Go"] = m["Go"] is equivalent to m["Go"] = 0.
After the assignment, the map contains one entry.
*/
func TestMap3(t *testing.T) {
	m := make(map[string]int, 3)
	x := len(m)
	m["Go"] = m["Go"]
	y := len(m)
	println(x, y)
}

// TODO:
func TestSyncMap(t *testing.T) {
	var wg sync.WaitGroup

	// ?
	var safeMap sync.Map
	safeMap.Store("k1", "v1")
	safeMap.Store("k2", "v2")
	v, ok := safeMap.Load("k1")
	ptr(v, ok)
	return
	wg.Add(2)
	// issue: fatal error: concurrent map writes
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			safeMap.Store("k1", "v1")
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 2000; i++ {
			// safeMap.Swap(key any, value any)
		}
	}()

	wg.Wait()
}

// TODO: advance.
func TestConcurrencyMap(t *testing.T) {
	var wg sync.WaitGroup

	builtinMap := make(map[string]int)

	wg.Add(2)
	// issue: fatal error: concurrent map writes
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			builtinMap["key"]++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 2000; i++ {
			builtinMap["key"]++
		}
	}()

	wg.Wait()
	ptr(builtinMap["key"])
}

// ?operation_test.go
// TODO: tidy.

// * bool operation priority problem.
// What does the following program print?
/*

When evaluating a || b, the expression a is evaluated firstly
and the expression b will be only evaluated if a is evaluated as false.

When evaluating a && b, the expression a is evaluated firstly
and the expression b will be only evaluated if a is evaluated as true.

*/
func TestBooleanPriority(t *testing.T) {
	o := func(b bool) bool {
		print(b)
		return !b
	}

	var x, y = true, false
	_ = x || o(x)
	_ = y && o(y)
}

// * byte is uint8 and overflow problem.
// What does the following program print?
/*

In Go, for a byte (a.k.a. uint8) non-constant non-zero value x,
-x overflows the range of type byte and is wrarpped as (-x + 256) % 256.

For a variable x of type byte, x == -x happens only when x is 0 or 128.

*/
func TestByteRangeWithOverflow(t *testing.T) {
	count := 0
	for i := range [256]struct{}{} {
		n := byte(i)
		a := n
		b := -n
		if a == b {
			count++
		}
	}
	ptr(count)
}

// *
// What does the following program print?
/*

1. If the operands in an operator expression are both/all constants,
then the expression is evaluated at compile time.
In the above program, `128 << N >> N` is such an expression.
n this expression, `128` is deduced as an untyped `int` value.

2. In a bit-shift expression, if the left operand is an untyped constant and the right operand is not constant,
then the type of the left operand will be deduced as the final assumed type.

In the above program, `128 << n >> n` is a such expression.
In this expression, the type of 128 is deduced as the assumed type, `byte`.

`128 << n` overflows the value range of `byte`, so it is truncated to 0.

*/
func TestOperandsWithConstant(t *testing.T) {
	const N = 1
	var n = N
	var a byte = 128 << N >> N
	var b byte = 128 << n >> n
	ptr(a, b)
}
