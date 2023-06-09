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
	ptr(e, &e, *e)
	ptr(e2, &e2, *e2, **e2)
	ptr(e3, &e3, *e3, **e3, ***e3)

	// send a ptr to function
	// v1 = 1, v2 = 2
	// &v1 = 0x1, &v2 = 0x2
	var v1, v2 = 1, 2
	changePtr(&v1)
	changeNonPtr(v2)
	ptr(v1, v2)
}

func changePtr(v *int) { // v = 0x1
	*v++
}

func changeNonPtr(v int) { // v = 2 , &v = 0x3
	v++
}

func TestPointerSize(t *testing.T) {
	var (
		user     User
		str      string
		num      int
		num8     int8
		num16    int16
		num32    int32
		num64    int64
		userPtr  *User
		strPtr   *string
		num32Ptr *int32
		num63Ptr *int64
	)
	ptr("User size:", unsafe.Sizeof(user), "byte")
	ptr("string size:", unsafe.Sizeof(str), "byte")
	ptr("int size:", unsafe.Sizeof(num), "byte")
	ptr("int8 size:", unsafe.Sizeof(num8), "byte")
	ptr("int16 size:", unsafe.Sizeof(num16), "byte")
	ptr("int32 size:", unsafe.Sizeof(num32), "byte")
	ptr("int64 size:", unsafe.Sizeof(num64), "byte")
	ptr("*User size:", unsafe.Sizeof(userPtr), "byte")
	ptr("*string size:", unsafe.Sizeof(strPtr), "byte")
	ptr("*int32 size:", unsafe.Sizeof(num32Ptr), "byte")
	ptr("*int64 size:", unsafe.Sizeof(num63Ptr), "byte")
}
