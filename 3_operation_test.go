package gotutorial

import (
	"testing"
)

func TestOperation(t *testing.T) {
	println("go" + "lang")

	println("1+1 =", 1+1)
	println("1-=1 =", 1-1)
	println("2*2 =", 2*2)
	println("7.0/3.0 =", 7.0/3.0)

	println(!true)

	println(true && false)
	println(true || false)

	println(10 | 4) // 1010 OR 0100 -> 1110 = 14
	println(3 & 6)  // 0011 OR 0110 -> 0010 = 2
}
