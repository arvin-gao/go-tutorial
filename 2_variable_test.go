package gotutorial

import (
	"errors"
	"math"
	"testing"
)

/*
#1.20 later
comparable #僅用在參數上使用(that can compare values with the == and != operators)
complex64 complex128

any #alias for interface{}
bool
error
string
float32 float64 #IEEE-754 32/64-bit
rune int int8 int16 int32 int64 #rune alias as int32
byte uint uint8 uint16 uint32 uint64 uintptr #byte alias as uint8
*/

/*
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
print(x*y)                 // "(-5+10i)"
print(real(x*y))           // "-5"
print(imag(x*y))           // "10"
*/
var (
	globalVar int = 10
)

const consGlobalVar = 10

func TestDeclaration(t *testing.T) {
	var (
		_a int
		_b int
	)

	var a int
	var b int = 1
	var c = 1
	var d *int
	e := 1

	var a1, b1, c1 = 1, 2.0, "3"

	v1, v2, err := getMultiVars()
	_, v22, _ := getMultiVars()

	const cInt = 1
	const cStr string = "1"

	pass(_a, _b, a, b, c, d, cInt, cStr, v1, v2, v22, err, a1, b1, c1, e)
}

func getMultiVars() (int, string, error) {
	return 0, "", errors.New("err")
}

func TestInitNil(t *testing.T) {
	var (
		ch       chan string
		m        map[string]interface{}
		in       interface{}
		num      *int
		user     *User
		slice    []int
		ptrSlice *[]int
	)
	ptr(
		ch == nil,
		m == nil,
		in == nil,
		num == nil,
		user == nil,
		slice == nil,
		ptrSlice == nil,
	)
}

// int8	  8  bits	-128 to 127
// int16  16 bits	-215 to 215 -1
// int32  32 bits	-231 to 231 -1
// int64  64 bits	-263 to 263 -1
// int depending on your system support(32 bit or 64 bit)

// uint8   8  bits	0 to 255
// uint16  16 bits	0 to 216 -1
// uint32  32 bits	0 to 232 -1
// uint64  64 bits	0 to 264 -1
// uint depending on your system support(32 bit or 64 bit)
func TestIntOverflow(t *testing.T) {
	var n int = math.MaxInt64
	ptr("n(MaxInt64) =", n)
	n = n + 1
	ptr("n + 1 =", n)
	ptr("n < 0 =", n < 0)
	n = n - 1
	ptr("n - 1 =", n)
	ptr("n < 0 =", n < 0)

	var n2 int32 = math.MaxInt32
	ptr("n2(MaxInt32) =", n2)
	n2 = n2 + 1
	ptr("n2 + 1 =", n2)
	ptr("n2 < 0 =", n2 < 0)
	var f float64 = math.MaxFloat64
	ptr("f(MaxFloat64) =", f)
	ptr("f + 1 =", f+1)
	ptr("f + f =", f+f)
	var f2 float32 = math.MaxFloat32
	ptr("f2(MaxFloat32) =", f2)
	ptr("f2 + 1 =", f2+1)
	ptr("f2 + f2 =", f2+f2)
}

// TODO: tidy.
/*
An interface value with a nil non-interface dynamic value is not a nil interface.
So the function f always return a non-nil interface value.

A call to the new builtin function always returns a pointer pointing to a zero value.
That means dereferencing the pointer results in a zero value.
The zero value of the *int type is nil.
*/
var x = *new(*int)
var y *int = nil

func f() interface{} {
	return y
}

func TestFuncNil(t *testing.T) {
	g := f()
	if g == nil {
		if x == nil {
			ptr("A")
		} else {
			ptr("B")
		}
	} else {
		if x == nil {
			ptr("C")
		} else {
			ptr("D")
		}
	}
}

// ----------------------------------------------------------------

/*
The nil in f(nil...) is treated as a nil slice (of type []interface{}).
The nil in f(nil) is treated as an element in a slice. That is equivalent to f([]interface{}{nil}...).
The call f() is equivalent to f(nil...).
*/
func TestInterfaceSlice(t *testing.T) {
	f := func(v ...interface{}) {
		ptr("val len:", len(v))
	}

	f()
	f(nil)
	f(nil...)
}
