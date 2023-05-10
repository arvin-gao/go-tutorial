package gotutorial

import (
	"testing"
)

/*
Defer is used to ensure that a function call is performed
later in a program’s execution, usually for purposes of cleanup.
*/
func TestDefer(t *testing.T) {

}

func TestDeferExample1(t *testing.T) {
	var a int
	a = 10
	defer println("example. ", a*a)
	a = 20
	deferB()
	a = 30
}

func deferB() {
	trace("b")
	defer unTrace("b")
	println("in b")
	deferA()
}

func deferA() {
	trace("a")
	defer unTrace("a")
	println("in a")
}

func trace(s string)   { println("entering:", s) }
func unTrace(s string) { println("leaving:", s) }

func TestDeferExample2(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", &a, calc("10", &a, b))
	a = 0
	defer calc("2", &a, calc("20", &a, b))
	b = 1
}

func calc(index string, a *int, b int) int {
	ret := *a + b
	println(index, *a, b, ret)
	return ret
}
