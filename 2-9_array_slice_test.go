package gotutorial

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// init
	var arr1 [2]int                             // [0,0]
	var arr2 = [2]int{}                         // [0,0]
	var arr3 = [2]int{1, 2}                     // [1,2]
	var arr4 = [...]int{1, 2, 2, 3, 3, 4, 4, 5} // [1,2]
	arr5 := [3][3]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	ptr(arr1, arr2, arr3, arr4, arr5)
	arr1[0] = 2
	// len cap
	fmt.Printf("var arr1 [2]int -> len:%d, cap:%d\n", len(arr1), cap(arr1))
}

func TestSlice(t *testing.T) {
	// init
	var slice1 []int                      // nil
	var slice2 = []int{}                  // []
	var slice3 = []int{1, 2}              // [1,2] -> len: 2, cap: 2
	slice3 = append(slice3, 2)            // -> len: 3, cap: 4
	slice3 = append(slice3, 2)            // -> len: 4, cap: 4
	slice3 = append(slice3, 2)            // -> len: 5, cap: 8
	var slice4 = make([]int, 2)           // [0,0], len:2, cap:2
	slice4 = append(slice4, 0)            // [0,0,0] len:3, cap:4
	var slice5 = make([]int, 2, 3)        // [0,0], len:2, cap:3
	slice5 = append(slice5, 0)            // [0,0,0] len:3, cap:3
	var slice6 = make([]int, 0, 3)        // []
	var slice7 = make([]int, len(slice3)) // [0,0]
	copy(slice7, slice3)
	slice8 := slice7[:len(slice7)-1]

	var ss = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	pCode("var ss = []int{0,1,2,3,4,5,6,7,8}")
	pCode("ss[:]")
	pfTree("%v", ss[:])
	pCode("ss[4:]")
	pfTree("%v", ss[4:])
	pCode("ss[:3]")
	pfTree("%v", ss[:3])
	pCode("ss[2:4]")
	pfTree("%v", ss[2:4])

	var ss2 []int
	ss2 = nil
	pCode("ss2 = nil")
	pfTree("ss2 ptr(%p)", ss2)
	ss2 = append(ss2, 1)
	pCode("ss2 = append(ss2, 1)")
	pfTree("ss2 ptr(%p)", ss2)
	// len cap
	pTitle("init")
	pCode("var slice1 []int")
	pSliceLenAndCap(slice1)
	pfTree("value: %v", slice1)
	pCode("var slice2 = []int{}")
	pSliceLenAndCap(slice2)
	pfTree("value: %v", slice2)
	pCode("var slice3 = []int{1, 2}")
	pSliceLenAndCap(slice3)
	pfTree("value: %v", slice3)
	pfTree("&[0]:%p", &slice3[0])
	pCode("var slice4 = make([]int, 2)")
	pSliceLenAndCap(slice4)
	pfTree("value: %v", slice4)
	pCode("var slice5 = make([]int, 2, 3)")
	pSliceLenAndCap(slice5)
	pfTree("value: %v", slice5)
	pCode("var slice6 = make([]int, 0, 3)")
	pSliceLenAndCap(slice6)
	pfTree("value: %v", slice6)
	pCode("var slice7 = make([]int, len(slice3))")
	pCode("copy(slice7, slice3)")
	pSliceLenAndCap(slice7)
	pfTree("value: %v", slice7)
	pfTree("&[0]:%p", &slice7[0])
	pCode("slice8 := slice7[:len(slice7)-1]")
	pSliceLenAndCap(slice8)
	pfTree("value: %v", slice8)
	pfTree("&[0]: %p", &slice8[0])

	pTitle("append()")
	// TODO: 檢查 append 後的slice地址是否相同？.
	s := make([]int, 2, 3)
	pCode("s := make([]int, 2, 3)")
	pfTree("s[0](%p)", &s[0])
	s = append(s, 1)
	pCode("s = append(s, 1)")
	pfTree("s[0](%p)", &s[0])
	s = append(s, 1)
	pCode("s = append(s, 1)")
	pfTree("s[0](%p)", &s[0])

	slice1 = append(slice1, []int{1, 2, 3}...)
	pCode("slice1 = append(slice1, []int{1,2,3}...)")
	pSliceLenAndCap(slice1)
	pfTree("value: %v", slice1)
	slice1 = append(slice1, []int{2, 3}...)
	pCode("slice1 = append(slice1, []int{2, 3}...)")
	pSliceLenAndCap(slice1)
	pfTree("value: %v", slice1)
	slice1 = append(slice1, []int{4}...)
	pCode("slice1 = append(slice1, []int{4}...)")
	pSliceLenAndCap(slice1)
	pfTree("value: %v", slice1)
	slice1 = append(slice1, []int{5, 6, 7}...)
	pCode("slice1 = append(slice1, []int{5,6,7}...)")
	pSliceLenAndCap(slice1)
	pfTree("value: %v", slice1)

	slice1 = make([]int, 1024)
	slice1 = append(slice1, []int{1}...)
	pCode("slice1 = make([]int, 1024)")
	pCode("slice1 = append(slice1, []int{1}...)")
	pSliceLenAndCap(slice1)
	// modify value
	pTitle("modify()")
	slice1 = make([]int, 5)
	slice1_by_directly := slice1[:]
	slice1_by_copy := make([]int, 5)
	copy(slice1_by_copy, slice1)
	slice1[0] = 1
	pCode("slice1 = make([]int, 5)")
	pCode("slice1_by_directly := slice1[:]")
	pCode("slice1_by_copy := make([]int, 5)")
	pCode("copy(slice1_by_copy, slice1)")
	pCode("slice1[0] = 1")
	pfTree("slice1[0]: %d (%p)", slice1[0], &slice1[0])
	pfTree("slice1_by_directly[0]: %d (%p)", slice1_by_directly[0], &slice1_by_directly[0])
	pfTree("slice1_by_copy[0]: %d (%p)", slice1_by_copy[0], &slice1_by_copy[0])
	slice1_by_directly = append(slice1_by_directly, 1)
	slice1[0] = 2
	pCode("slice1_by_directly = append(slice1_by_directly, 1)")
	pCode("slice1[0] = 2")
	pfTree("slice1[0]: %d (%p)", slice1[0], &slice1[0])
	pfTree("slice1_by_directly[0]: %d (%p)", slice1_by_directly[0], &slice1_by_directly[0])

	changeFirstValue := func(s []int) {
		s[0] = 100
	}
	slice1[0] = 1
	pCode("slice1[0] = 1")
	pfTree("slice1[0]: %d (%p)", slice1[0], &slice1[0])
	changeFirstValue(slice1)
	pCode("changeFirstValue(slice1) // func(s []int) { s[0] = 100 }")
	pfTree("slice1[0]: %d (%p)", slice1[0], &slice1[0])
}

func TestSliceExample1(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	ptr(s1)
	ptr(s2)
	ptr(slice)
}

func TestSliceExample2(t *testing.T) {
	changeFirstBeforeAppend := func(s []int) {
		s = append(s, 0)
		s[0] = 102
	}

	changeFirst := func(s []int) {
		s[0] = 101
	}

	slice := make([]int, 5)
	ptr("origin slice[0]: ", slice[0])

	changeFirstBeforeAppend(slice)
	ptr("changeFirstBeforeAppend slice[0]: ", slice[0])

	changeFirst(slice)
	ptr("changeFirst slice[0]: ", slice[0])
}

func TestSliceExample3(t *testing.T) {
	myAppend := func(s []int) []int {
		s = append(s, 100)
		return s
	}
	s := []int{1, 1, 1}
	newS := myAppend(s)

	ptr(s)
	ptr(newS)

	s = newS

	myAppendPtr := func(s *[]int) {
		*s = append(*s, 100)
	}
	myAppendPtr(&s)
	ptr(s)
}

func TestSliceExample4(t *testing.T) {
	s := []int{1}
	s = append(s, 2)
	pCode("s := []int{1}")
	pCode("s = append(s, 2)")
	pPtr(s)

	s = append(s, 3)
	pCode("s = append(s, 3)")
	pPtr(s)

	x := append(s, 11)
	pCode("x := append(s, 11)")
	pfTree("x = %v", x)

	y := append(s, 22)
	pCode("y := append(s, 22)")
	pfTree("y = %v", y)

	z := []int{2: 5, 6, 0: 7}
	pCode("z := []int{2: 5, 6, 0: 7}")
	pfTree("z = %v", z)

	y[len(y)-1] = 23
	pCode("y[len(y)-1] = 23")
	pfTree("s = %v", s)
	pfTree("x = %v", x)
	pfTree("y = %v", y)
	pCode("s[0] = 0")
	s[0] = 0
	print("s: ")
	pPtr(s)
	pCode("s = append(s, 4)")
	s = append(s, 4)
	print("s: ")
	pPtr(s)
	ptr("s:", s)
	ptr("x:", x)
	ptr("y:", y)
	pCode("s = append(s, 5)")
	s = append(s, 5)
	print("s: ")
	pPtr(s)
	// 總結：當 slice 容量還夠時，slice 的 struct 及 array 地址並不會變
}

func TestSliceExample5(t *testing.T) {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	x = append(x, y...)
	x = append(x, y...)
	ptr(a, x)
}

func TestSliceExample6(t *testing.T) {
	var x = []string{"A", "B", "C"}

	for i, s := range x {
		print(i, s, ",")
		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}
