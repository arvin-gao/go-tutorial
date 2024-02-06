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
func TestBasicDefer(t *testing.T) {
	executePanic()
	ptr("done")
}

func executePanic() {
	// Defer function call
	defer recoveryFunction()
	panic("Panic")
	ptr("executePanic function done")
}

func recoveryFunction() {
	// Recover from panic and print erro message
	if recoveryMessage := recover(); recoveryMessage != nil {
		ptr(recoveryMessage)
	}
	ptr("Recovery function done")
}

func TestDeferExample1(t *testing.T) {
	enter("example")
	defer leave("example")
	deferA()
}

func deferA() {
	enter("a")
	defer leave("a")
	ptr("process a")
	deferB()
}

func deferB() {
	enter("b")
	defer leave("b")
	ptr("process b")
}

func enter(s string) { ptr("entering:", s) }
func leave(s string) { ptr("leaving:", s) }

func TestDeferExample2(t *testing.T) {
	calc := func(index string, a *int, b int) int {
		ret := *a + b
		fmt.Printf("[%s]: a=%d, b=%d, a+b=%d\n", index, *a, b, ret)
		return ret
	}

	a := 1
	b := 2
	defer calc("1-1", &a, calc("1-2", &a, b))
	ptr("step-1")
	a = 2
	b = 3
	defer calc("2-1", &a, calc("2-2", &a, b))
	a = 3
	b = 4
	ptr("step-2")
}
