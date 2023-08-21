package gotutorial

import (
	"fmt"
	"testing"
)

func TestBasicTraining(t *testing.T) {
	// * Question 1: How to check type of variable by fmt package?
	// * Question 2: How to count of UTF-8 string and what length of '你' and 'é' word?
	// * Question 3: How to use json.NewEncoder with bytes.Buffer and does marshaled string is equal to buffer string?
	// * Question 4: How to compare with map and slices?
}

func TestInterfaceTraining(t *testing.T) {
	ptr("answers:")
	// * Question 1
	var data *byte
	var in interface{}
	// ? What result?
	ptr(
		data == nil,
		in == nil,
	)
	in = data
	// ? What result?
	ptr(in == nil)

	// * Question 2
	f := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		}
		return result
	}
	if res := f(-1); res != nil {
		// ? What result?
		ptr("result:", res)
	}
}

func TestDeferTraining(t *testing.T) {
	// * Question 1
	for i := 0; i < 5; i++ {
		// ? What result?
		defer ptr(i)
	}

	// * Question 2
	un := func(s string) {
		println("leaving:", s)
	}

	trace := func(s string) string {
		println("entering:", s)
		return s
	}
	a := func() {
		defer un(trace("a"))
		println("in a")
	}
	b := func() {
		defer un(trace("b"))
		println("in b")
		a()
	}
	// ? What result?
	b()
}

func doRecover() {
	fmt.Println("print:", recover())
}

func TestDeferWithPanic1Training(t *testing.T) {
	defer func() {
		doRecover()
	}()
	// ? What result?
	panic("test")
}
func TestDeferWithPanic2Training(t *testing.T) {
	// ? What result?
	defer doRecover() // print: test
	panic("test")
}
func TestDeferWithPanic3Training(t *testing.T) {
	defer func() {
		// ? What result?
		ptr(recover())
	}()
	defer func() {
		// ? What result?
		ptr(recover())
	}()
	panic(1)
}

func TestMoreDeferTraining(t *testing.T) {
	i := 1
	println("begin i =", i)

	defer func() {
		// ? What result?
		ptr("result 9:", i*2)
	}()
	defer func(i2 int) {
		// ? What result?
		ptr("result 8:", i2*2)
	}(i)
	defer func(i2 *int) {
		// ? What result?
		ptr("result 7:", *i2*2)
	}(&i)
	defer func(i2 int) {
		// ? What result?
		ptr("result 6:", func(i3 int) int {
			return i3 * 2
		}(i2))
	}(i)
	defer func(i2 int) {
		// ? What result?
		ptr("result 5:", func(i3 *int) int {
			return *i3 * 2
		}(&i2))
	}(i)
	defer func(i2 *int) {
		// ? What result?
		ptr("result 4:", func(i3 *int) int {
			return *i3 * 2
		}(&i))
	}(&i)
	// ? What result?
	defer ptr("result 3:", i*2)
	// ? What result?
	defer ptr("result 2:", func() int {
		return i * 2
	}())
	// ? What result?
	defer ptr("result 1:", func(i2 *int) int {
		return *i2 * 2
	}(&i))

	i = 2
	println("end i =", i)
}

type user struct {
	name string
}

func (u *user) print() {
	ptr(u.name)
}

func TestGroutineTraining(t *testing.T) {
	// * Question 1: What happens with the code below?
	// ? Answer: ?
	// go fmt.Println("3s", <-time.Tick(3*time.Second))
	// println("done")
	// * Question 2: What happens with the code below?
	// ? Answer: ?
	// var c chan int
	// go func() { c <- 1 }()
	// <-c
	// * Question 3
	data := []*user{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		// ? What result?
		go v.print()
		// ? What result?
		go func() { v.print() }()
	}
	data2 := []user{{"one"}, {"two"}, {"three"}}
	for _, v := range data2 {
		// ? What result?
		go v.print()
		// ? What result?
		go func() { v.print() }()
	}
}

func TestLoopTraining(t *testing.T) {
	// *Question 1
	data := []string{"1", "2", "3"}
	for _, v := range data {
		go func() {
			// ? What result?
			println(v)
		}()
	}
	// *Question 2: What happens when running the code below?
	// ? Answer: ?
	// for i := 0; i < len(data); i++ {
	// 	go func() {
	// 		println(data[i])
	// 	}()
	// }
	// *Question 3: How to resolve the question 1 problem?
}

func TestOtherTraining(t *testing.T) {
	// * Question 1: What statements different `break`, `return`, `goto out`, `break out`.
out:
	for {
		switch {
		case true:
			println("breaking out...")
			// break
			// return
			// goto out
			break out
		}
	}
	println("out")

	// * Question 2: What happens to running the code below and how to resolves the problem?
	var v1 interface{} = "s"
	if v1, ok := v1.(int); ok {
	} else {
		// ? What result?
		println(v1)
	}
	// ? Solution
}

type dataWithReciver struct {
	name string
}

func (p *dataWithReciver) print() { println(p.name) }

type dataWithoutReciver struct {
	name string
}

func (p dataWithoutReciver) print() { println(p.name) }

type printer interface {
	print()
}

func TestAddressTraining(t *testing.T) {
	// * Question 1: What happens the code below?
	// var _ printer = dataWithReciver{}
	// var _ printer = dataWithoutReciver{}
	// var _ printer = &dataWithReciver{}
	// Map with a struct with reciver function.

	// * Question 2
	m1 := map[string]dataWithoutReciver{"key": {"a"}}
	m1["key"].print()
	// * What happens the code?
	// m2 := map[string]dataWithReciver{"key": dataWithReciver{"a"}}
	// m2["key"].print()

	// * What happens the code?
	// m1["key"].name = "b"

	v := m1["key"]
	v.name = "b"
	m1["key"] = v

	m3 := map[string]*user{"key": {"a"}}
	m3["key"].name = "b"
	// * What happens the code?
	// m3["notExistKey"].name = "a"
}
