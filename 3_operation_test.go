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
