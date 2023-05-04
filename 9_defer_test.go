package gotutorial

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {

}

func TestDeferExample1(t *testing.T) {
	var a int
	a = 10
	defer fmt.Println("example. ", a*a)
	a = 20
	deferB()
	a = 30
}

func deferB() {
	trace("b")
	defer unTrace("b")
	fmt.Println("in b")
	deferA()
}

func deferA() {
	trace("a")
	defer unTrace("a")
	fmt.Println("in a")
}

func trace(s string)   { fmt.Println("entering:", s) }
func unTrace(s string) { fmt.Println("leaving:", s) }

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
	fmt.Println(index, *a, b, ret)
	return ret
}
