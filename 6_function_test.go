package gotutorial

import (
	"errors"
	"testing"
)

func TestFunction(t *testing.T) {
	// Call simple functions
	simpleFunc(1, 2, "v3 string", "str2")

	// Passing multiple values using a variadic function.
	slice := []string{"str1", "str2"}
	simpleFunc(1, 2, "v3", slice...)
	simpleFunc(1, 2, "v4", []string{"str3", "str4"}...)
	simpleFunc(1, 2, "v4", "str4", "str5")

	// Function return multiple values.
	v1, v2 := simpleFunc2()
	ptr(v1, v2)

	// Anonymous function
	func() {
		ptr("Anonymous function")
	}()

	func1 := func(x int) int {
		return x + 1
	}
	ptr(func1(2))

	vByFunc := func(x int) int {
		return x + 1
	}(3)
	ptr(vByFunc)

	// The function value as param of function call.
	printFWithTen(func(x int) int {
		return x + 123
	})

	// Closure function(閉包)
	pTitle("Closure function")
	f := closureFunction()
	ptr("f(1):", f(1))
	ptr("f(2):", f(2))

	// Closure function 2
	l := 10
	b := 10
	func() {
		var area int
		area = l * b
		ptr(area)
	}()

	// 遞迴
	pTitle("Recursion")
	ptr("fib(7):", fib(3))
}

func TestFunctionWithCondition(t *testing.T) {
	f := func() error {
		return errors.New("testing error")
	}
	if err := f(); err != nil {
		panic(err)
	}
}

func closureFunction() func(b int) int {
	var num int
	ptrlnf("num(outside) address: %p", &num)
	return func(n int) int {
		num++
		ptrlnf("num(inside) address:%p. value: %d", &num, num)
		return num + n
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
	z := v3 + "hi"

	for _, v := range strSlice {
		ptr(v)
	}

	return x, y, z
}

func simpleFunc2() (x, y int) {
	x = 1
	y = 2
	return // 也可直接使用 return 1, 2
}

func TestAliasFunction(t *testing.T) {
	var p = ptr

	myFunc := func() {
		p("my function")
	}

	var f = myFunc
	f()
}
