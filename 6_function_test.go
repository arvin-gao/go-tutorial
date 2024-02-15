package gotutorial

import (
	"errors"
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	// Call simple functions
	simpleFunc(1, 2, "one", "s1")

	// Passing multiple values using a variadic function.
	slice := []string{"s1", "s2"}
	simpleFunc(1, 2, "two", slice...)
	simpleFunc(1, 2, "three", []string{"s1", "s2"}...)
	simpleFunc(1, 2, "four", "s1", "s2")

	// Function return multiple values.
	_, _ = returnDefinedVariable()

	// Anonymous function
	func() {
		// Do something there.
	}()

	// Define a function to the variable.
	vF1 := func(x int) int {
		return x + 1
	}
	ptr(vF1(10))

	// Call a anonymous function and return a value to the variable.
	v := func(x int) int {
		return x + 1
	}(3)
	ptr(v)

	// The function value as param of function call.
	printFWithTen(func(x int) int {
		return x + 123
	})

	// Closure function(閉包)
	l := 10
	b := 10
	func() {
		var area int
		area = l * b
		ptr(area)
	}()

	// Closure function-2
	closureFunction := func() func(b int) {
		num := 2
		ptrlnf("num(outside) address:%p. value: %d", &num, num)
		return func(n int) {
			num += n
			ptrlnf("num(inside)  address:%p. value: %d", &num, num)
		}
	}
	closureFunction2 := closureFunction()
	closureFunction2(1) // 3
	closureFunction2(1) // 4
	closureFunction2(2) // 6

	// Recursion from function
	_ = fib(3)

	// Recursion from Closure function
	var fib2 func(n int) int
	fib2 = func(n int) int {
		if n < 2 {
			return n
		}
		return fib2(n-1) + fib2(n-2)
	}
	_ = fib2(10)
}

func TestFunctionWithCondition(t *testing.T) {
	f := func() error {
		return errors.New("testing error")
	}
	if err := f(); err != nil {
		panic(err)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func printFWithTen(f func(int) int) {
	ptr(f(10))
}

func simpleFunc(v, v2 int, v3 string, strSlice ...string) (int, int, string) {
	x := v + 1
	y := v2 + 2
	z := v3

	fmt.Print(v3, ":")
	for _, v := range strSlice {
		fmt.Print(" ", v)
	}
	ptr()
	return x, y, z
}

func returnDefinedVariable() (x, y int) {
	x = 1
	y = 2
	return // 也可直接使用 return 1, 2
}

func TestAliasFunction(t *testing.T) {
	var p = ptr

	f := func() {
		p("f")
	}

	var f2 = f
	f2()
}
