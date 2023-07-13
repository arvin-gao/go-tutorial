package gotutorial

import (
	"fmt"
	"testing"
)

/*
Defer is used to ensure that a function call is performed
later in a program’s execution, usually for purposes of cleanup.
*/
// FILO. First in last out.

func TestDeferExample1(t *testing.T) {
	enter("example-1")
	defer leave("example-1")
	deferB()
}

func deferB() {
	enter("b")
	defer leave("b")
	ptr("process b")
	deferA()
}

func deferA() {
	enter("a")
	defer leave("a")
	ptr("process a")
}

func enter(s string) { ptr("entering:", s) }
func leave(s string) { ptr("leaving:", s) }

func TestDeferExample2(t *testing.T) {
	a := 1
	b := 2
	defer calc("1-1", &a, calc("1-2", &a, b))
	ptr("[0-1]")
	a = 2
	b = 3
	defer calc("2-1", &a, calc("2-2", &a, b))
	a = 3
	b = 4
	ptr("[0-2]")
}

func calc(index string, a *int, b int) int {
	ret := *a + b
	fmt.Printf("[%s]: a=%d, b=%d, a+b=%d\n", index, *a, b, ret)
	return ret
}
