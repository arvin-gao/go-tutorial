package gotutorial

import (
	"testing"
)

func TestFunction(t *testing.T) {
	pTitle("call simple functions ")

	simpleFunc(1, 2, "v3 string", "str2")

	slice := []string{"str1", "str2"}
	simpleFunc(1, 2, "v3", slice...)

	v1, v2 := simpleFunc2()
	ptr(v1, v2)
	ptr(simpleFunc2())

	pTitle("匿名函式")
	func() {
		ptr("1")
	}()

	func1 := func(x int) int {
		return x + 1
	}
	ptr(func1(2))

	vByFunc := func(x int) int {
		return x + 1
	}(3)
	ptr(vByFunc)

	pTitle("function 當作參數傳遞")
	sendFunc(func(x int) int {
		ptr("this is my x:", x)
		return x + 123
	})

	// 閉包
	f := closures()
	ptr("f(1):", f(1))
	ptr("f(2):", f(2))

	// 遞迴
	ptr("fib(7):", fib(3))
}

func closures() func(b int) int {
	var num int
	ptr("closure init, num address:", &num)
	return func(n int) int {
		num++
		pf("num(%p): %d", &num, num)
		return num + n
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func sendFunc(f func(int) int) {
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
