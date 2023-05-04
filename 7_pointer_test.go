package gotutorial

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {
	var e *int = new(int)
	var e2 **int = &e
	var e3 ***int = &e2
	*e = 1
	fmt.Println(e, &e, *e)
	fmt.Println(e2, &e2, *e2, **e2)
	fmt.Println(e3, &e3, *e3, **e3, ***e3)

	// send a ptr to function
	var v1, v2 = 1, 2
	changePtr(&v1)
	changeNonPtr(v2)
	fmt.Println(v1, v2)
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
	fmt.Println("User:", unsafe.Sizeof(u1))
	fmt.Println("string:", unsafe.Sizeof(s1))
	fmt.Println("int64:", unsafe.Sizeof(n1))
	fmt.Println("int32:", unsafe.Sizeof(n11))
	fmt.Println("*User:", unsafe.Sizeof(u2))
	fmt.Println("*string:", unsafe.Sizeof(s2))
	fmt.Println("*int64:", unsafe.Sizeof(n2))
	fmt.Println("*int32:", unsafe.Sizeof(n21))
}
