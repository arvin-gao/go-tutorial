package gotutorial

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	pTitle("call simple functions ")
	simpleFunc(1, 2, "str1", "str2")
	v1, v2 := simpleFunc2()
	println(v1, v2)
	println(simpleFunc2())

	pTitle("匿名函式")
	func() {
		println("1")
	}()

	a := func(x int) int {
		return x + 1
	}
	println(a(2))

	b := func(x int) int {
		return x + 1
	}(3)
	println(b)

	pTitle("function 當作參數傳遞")
	sendFunc(func(x int) int {
		fmt.Println("this is my x:", x)
		return x + 123
	})

	// 閉包
	f := Closure()
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func Closure() func(b int) int {
	var num int
	return func(b int) int {
		num++
		fmt.Println("num:", num)
		return b + 2
	}
}

func sendFunc(f func(int) int) {
	println(f)
	fmt.Println(f(10))
}

func simpleFunc(v, v2 int, strSlice ...string) (int, int) {
	x := v + 1
	y := v + 2

	for _, v := range strSlice {
		println(v)
	}

	return x, y
}

func simpleFunc2() (x, y int) {
	x = 1
	y = 2
	return // 也可直接使用 return 1, 2
}
