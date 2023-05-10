package gotutorial

import (
	"testing"
)

func TestOperation(t *testing.T) {
	println("go" + "lang")

	println("1+1 =", 1+1)
	println("7.0/3.0 =", 7.0/3.0)

	println(true && false)
	println(true || false)
	println(!true)
}
