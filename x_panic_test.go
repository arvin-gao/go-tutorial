package gotutorial

import (
	"fmt"
	"testing"
)

/*
A panic typically means something went unexpectedly wrong

A common use of panic is to abort if a function returns an
error value that we don’t know how to (or want to) handle.
*/
func TestPanic(t *testing.T) {
	panic("panic message")
}

/*
Recover can stop a panic from aborting the program and let
it continue with execution instead

# Recover must be called within a deferred function

An example of where this can be useful:
A server wouldn’t want to crash if one of the client connections
exhibits a critical error. Instead, the server would want to
close that connection and continue serving other clients.
*/
func TestPanicWithRecover(t *testing.T) {
	defer func() {
		println("done 2")
		if v := recover(); v != nil {
			println("recover 2:", v)
		}
	}()

	println("start")

	defer func() {
		println("done 1")
		if v := recover(); v != nil {
			println("recover 1:", v)
		}
	}()

	panic("something panic!")
}

/*
1. The two recover calls at comment 1 and 2 are no-op.
2. The recover calls at comment 3 catches the panic 2.
3. The recover calls at comment 4 catches the panic 1.
？ [Explain Panic/Recover Mechanism in Detail](https://go101.org/article/panic-and-recover-more.html)
*/
func TestDeferAndRecoverMechanism(t *testing.T) {
	defer func() {
		fmt.Print(recover()) // 4
	}()

	defer func() {
		defer fmt.Print(recover()) // 3
		defer panic(1)
		recover() // 2
	}()

	defer recover() // 1
	panic(2)
}

// *

/*
Except the two recover calls at comment 2 and 1, the other ones are all no-op.
The recover calls at comment 2 recovers the panic 3.
The recover calls at comment 1 recovers the panic 2.
The the panic 1 is never recovered, but it is suppressed by the panic 2.
*/
func TestMoreReferAndRecover(t *testing.T) {
	defer func() {
		println(recover().(int)) // 1
	}()

	defer func() {
		defer func() {
			recover() // 2
		}()
		defer recover()
		panic(3)
	}()

	defer func() {
		defer func() {
			defer func() {
				recover()
			}()
		}()
		defer recover()
		panic(2)
	}()

	panic(1)
}
