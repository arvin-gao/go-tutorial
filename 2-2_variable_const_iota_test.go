package gotutorial

import (
	"testing"
)

const (
	_x = iota // 0
	_y = iota // 1
	_z = iota // 2
)

const (
	_a = iota     // 0
	_b            // 1
	_c            // 2
	_d = 5        // 5
	_e            // 5
	_f            // 5
	_g = iota + 2 // 8
	_h            // 9
)

func Test(t *testing.T) {
	ptr(_a)
	ptr(_b)
	ptr(_c)
	ptr(_d)
	ptr(_e)
	ptr(_f)
	ptr(_g)
	ptr(_h)
}

const (
	Apple, Banana     = iota + 1, iota + 2 // 1, 2
	Cherimoya, Durian                      // 2, 3
	Elderberry, Fig                        // 3, 4

)

const (
	Open    = 1 << iota // 0001(1)
	Close               // 0010(2)
	Pending             // 0100(4)
)

const (
	_  = iota             // 使用 _ 忽略不需要的 iota
	KB = 1 << (10 * iota) // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
	TB                    // 1 << (10*4)
	PB                    // 1 << (10*5)
	EB                    // 1 << (10*6)
	ZB                    // 1 << (10*7)
	YB                    // 1 << (10*8)
)

const (
	_k = 3 // now, iota == 0

	_m float32 = iota + .5 // m float32 = 1 + .5
	_n                     // n float32 = 2 + .5

	_p     = 9          // now, iota == 3
	_q     = iota * 2   // q = 4 * 2
	_                   // _ = 5 * 2
	_r                  // r = 6 * 2
	_s, _t = iota, iota // s, t = 7, 7
	_u, _v              // u, v = 8, 8
	_, _w               // _, w = 9, 9
)

// Untyped constant can overflow its default type
const _overflowInt = 1 << 64           // overflows int
const _overflowRune = 'a' + 0x7FFFFFFF // overflows rune
const _overflowFloat = 2e+308          // overflows float64

// TODO: tidy
/*
Go spec says: Within a parenthesized const declaration list
the expression list may be omitted from any but the first ConstSpec.

Such an empty list is equivalent to the textual substitution of
the first preceding non-empty expression list and its type if any.
*/

// * iota
/*
The value of the predeclared iota is the constant specification order id (0-based) in a constant declaration.
*/
func TestConstantWithIota(t *testing.T) {
	const (
		_    = 6
		A, _ = iota, iota + 10 // 1, 1 + 10
		_, _                   // 2, 2 + 10
		_, B                   // 3, 3 + 10
	)
	ptr(A, B)
}

// * const mechanism
/*
A local identifier will shadow the global identifier with the same name.

! Notice: The output result was 6 6 when using Go toolchain v1.17-. The bug has been fixed since Go toochain v1.18.
*/

const X = 3

func TestConstantWithGlo(t *testing.T) {
	const (
		X = X + X // X + X. The two "X" are both the `global` one
		Y         // X + X. The two "X" are both the `local` one
	)

	ptr(X, Y)
}

// * ???

/*
Untyped X is a rune constant (in other words, its default type is rune, a.k.a int32, a 32-bit integer type).
Untyped Y is an int constant (in other words, its default type is int, a 64-bit integer type on 64-bit OSes).

At run time, the expression A << n overflows, so it is evaluated as 0; on the other hand, the expression B << n doesn't overflow.
*/
func TestConstantWithOperands(t *testing.T) {
	const (
		X       = '\x61' // 'a'
		Y       = 0x62
		A       = Y - X // 1
		B int64 = 1
	)

	var n = 32
	if A == B {
		ptr(A<<n>>n, B<<n>>n)
	}
}
