package gotutorial

import (
	"testing"
)

func TestOperation(t *testing.T) {
	ptr("\"go\" + \"lang\" = ", "go"+"lang")

	ptr("1+1 =", 1+1)
	ptr("1-1 =", 1-1)
	ptr("2*2 =", 2*2)
	ptr("7%3 =", 7%3)
	ptr("7.0/3.0 =", 7.0/3.0)

	ptr("!true:", !true)
	ptr("true && false:", true && false)
	ptr("true || false:", true || false)

	ptr("10 | 4 =", 10|4) // 1010 OR 0100 -> 1110 = 14
	ptr("3 & 6 =", 3&6)   // 0011 OR 0110 -> 0010 = 2

	pTitle("比較")
	ptrf("10 >  5: %t", 10 > 5)
	ptrf("10 >= 5: %t", 10 >= 5)
	ptrf("10 <  5: %t", 10 < 5)
	ptrf("10 <= 5: %t", 10 <= 5)
	ptrf("10 == 5: %t", 10 == 5)
	ptrf("10 != 5: %t", 10 != 5)

	pTitle("AND")
	ptrf("0&1: %2d", 0&1)
	ptrf("0&1: %2d", 0&1)
	ptrf("1&0: %2d", 1&0)
	ptrf("1&1: %2d", 1&1)

	pTitle("OR")
	ptrf("0|0: %2d", 0|0)
	ptrf("0|1: %2d", 0|1)
	ptrf("1|0: %2d", 1|0)
	ptrf("1|1: %2d", 1|1)

	pTitle("XOR")
	ptrf("0^0: %2d", 0^0)
	ptrf("0^1: %2d", 0^1)
	ptrf("1^0: %2d", 1^0)
	ptrf("1^1: %2d", 1^1)

	pTitle("AND NOT")
	ptrf("0&^0: %1d", 0&^0)
	ptrf("0&^1: %1d", 0&^1)
	ptrf("1&^0: %1d", 1&^0)
	ptrf("1&^1: %1d", 1&^1)

	pTitle("位移")
	ptrf("1<<1: %d", 1<<1)
	ptrf("1<<2: %d", 1<<2)
	ptrf("1<<3: %d", 1<<3)
}

// TODO: tidy.

// * bool operation priority problem.
// What does the following program print?
/*

When evaluating a || b, the expression a is evaluated firstly
and the expression b will be only evaluated if a is evaluated as false.

When evaluating a && b, the expression a is evaluated firstly
and the expression b will be only evaluated if a is evaluated as true.

*/
func TestBooleanPriority(t *testing.T) {
	o := func(b bool) bool {
		print(b)
		return !b
	}

	var x, y = true, false
	_ = x || o(x)
	_ = y && o(y)
}

// * byte is uint8 and overflow problem.
// What does the following program print?
/*

In Go, for a byte (a.k.a. uint8) non-constant non-zero value x,
-x overflows the range of type byte and is wrarpped as (-x + 256) % 256.

For a variable x of type byte, x == -x happens only when x is 0 or 128.

*/
func TestByteRangeWithOverflow(t *testing.T) {
	count := 0
	for i := range [256]struct{}{} {
		n := byte(i)
		a := n
		b := -n
		if a == b {
			count++
		}
	}
	ptr(count)
}

// *
// What does the following program print?
/*

1. If the operands in an operator expression are both/all constants,
then the expression is evaluated at compile time.
In the above program, `128 << N >> N` is such an expression.
n this expression, `128` is deduced as an untyped `int` value.

2. In a bit-shift expression, if the left operand is an untyped constant and the right operand is not constant,
then the type of the left operand will be deduced as the final assumed type.

In the above program, `128 << n >> n` is a such expression.
In this expression, the type of 128 is deduced as the assumed type, `byte`.

`128 << n` overflows the value range of `byte`, so it is truncated to 0.

*/
func TestOperandsWithConstant(t *testing.T) {
	const N = 1
	var n = N
	var a byte = 128 << N >> N
	var b byte = 128 << n >> n
	ptr(a, b)
}
