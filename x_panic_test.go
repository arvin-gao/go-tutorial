package gotutorial

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	defer func() {
		println("done 2")
		if v := recover(); v != nil {
			fmt.Println("recover 2:", v)
		}
	}()

	println("start")

	defer func() {
		println("done 1")
		if v := recover(); v != nil {
			fmt.Println("recover 1:", v)
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
