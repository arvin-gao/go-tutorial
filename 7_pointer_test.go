package gotutorial

import (
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {
	var e *int = new(int)
	var e2 **int = &e
	var e3 ***int = &e2
	*e = 1
	println(e, &e, *e)
	println(e2, &e2, *e2, **e2)
	println(e3, &e3, *e3, **e3, ***e3)

	// send a ptr to function
	var v1, v2 = 1, 2
	changePtr(&v1)
	changeNonPtr(v2)
	println(v1, v2)
}

func changePtr(v *int) {
	*v++
}

func changeNonPtr(v int) {
	v++
}

func TestPointerSize(t *testing.T) {
	var (
		u1  User
		s1  string
		n1  int64
		n11 int32
		u2  *User
		s2  *string
		n2  *int64
		n21 *int32
	)
	println("User size:", unsafe.Sizeof(u1))
	println("string size:", unsafe.Sizeof(s1))
	println("int64 size:", unsafe.Sizeof(n1))
	println("int32 size:", unsafe.Sizeof(n11))
	println("*User size:", unsafe.Sizeof(u2))
	println("*string size:", unsafe.Sizeof(s2))
	println("*int64 size:", unsafe.Sizeof(n2))
	println("*int32 size:", unsafe.Sizeof(n21))
	// TODO: compare size of param with the pointer and normal param.
}
