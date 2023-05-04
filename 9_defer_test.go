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
