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
	f := func(s string, sizeOf uintptr) string { return fmt.Sprintf("%s size: %v byte", s, sizeOf) }
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
	str2 := "xxx"
	ptrs(
		f("User", unsafe.Sizeof(user)),              // 24 byte.
		f("*User", unsafe.Sizeof(userPtr)),          // 8 byte.
		f("string(empty)", unsafe.Sizeof(str)),      // 16 byte.
		f("string(with text)", unsafe.Sizeof(str2)), // 16 byte.
		f("*string", unsafe.Sizeof(strPtr)),         // 8 byte.
		f("int", unsafe.Sizeof(num)),                // 8 byte.
		f("int8", unsafe.Sizeof(num8)),              // 1 byte.
		f("int16", unsafe.Sizeof(num16)),            // 2 byte.
		f("int32", unsafe.Sizeof(num32)),            // 4 byte.
		f("*int32", unsafe.Sizeof(num32Ptr)),        // 8 byte.
		f("int64", unsafe.Sizeof(num64)),            // 8 byte.
		f("*int64", unsafe.Sizeof(num63Ptr)),        // 8 byte.
	)
}
