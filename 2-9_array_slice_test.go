package gotutorial

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// init
	var arr1 [2]int           // [0,0]
	var arr2 = [2]int{}       // [0,0]
	var arr3 = [2]int{1, 2}   // [1,2]
	var arr4 = [...]int{1, 2} // [1,2]
	arr5 := [3][3]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	println(arr1, arr2, arr3, arr4, arr5)
	// len cap
	fmt.Printf("var arr1 [2]int -> len:%d, cap:%d\n", len(arr1), cap(arr1))
}

func TestSlice(t *testing.T) {
	// init
	var slice1 []int // []
	// var slice2 = []int{}                  // []
	var slice3 = []int{1, 2} // [1,2]
	// var slice4 = make([]int, 2)           // [0,0]
	// var slice5 = make([]int, 2, 3)        // [0,0]
	// var slice6 = make([]int, 0, 3)        // []
	var slice7 = make([]int, len(slice3)) // [0,0]
	copy(slice7, slice3)
	// slice8 := slice7[:len(slice7)-1]

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
	// printfCode("var slice1 []int")
	// printSliceLenAndCap(slice1)
	// printfChild("value: %v", slice1)
	// printfCode("var slice2 = []int{}")
	// printSliceLenAndCap(slice2)
	// printfChild("value: %v", slice2)
	// printfCode("var slice3 = []int{1, 2}")
	// printSliceLenAndCap(slice3)
	// printfChild("value: %v", slice3)
	// printfChild("&[0]:%p", &slice3[0])
	// printfCode("var slice4 = make([]int, 2)")
	// printSliceLenAndCap(slice4)
	// printfChild("value: %v", slice4)
	// printfCode("var slice5 = make([]int, 2, 3)")
	// printSliceLenAndCap(slice5)
	// printfChild("value: %v", slice5)
	// printfCode("var slice6 = make([]int, 0, 3)")
	// printSliceLenAndCap(slice6)
	// printfChild("value: %v", slice6)
	// printfCode("var slice7 = make([]int, len(slice3))")
	// printfCode("copy(slice7, slice3)")
	// printSliceLenAndCap(slice7)
	// printfChild("value: %v", slice7)
	// printfChild("&[0]:%p", &slice7[0])
	// printfCode("slice8 := slice7[:len(slice7)-1]")
	// printSliceLenAndCap(slice8)
	// printfChild("value: %v", slice8)

	pTitle("append()")
	// TODO: 檢查 append 後的slice地址是否相同？.
	s := make([]int, 2, 3)
	pCode("s := make([]int, 2, 3)")
	pfTree("s address(%p)", s)
	pfTree("s[0](%p)", s[0])
	s = append(s, 1)
	pCode("s = append(s, 1)")
	pfTree("s address(%p)", s)
	pfTree("s[0](%p)", s[0])
	s = append(s, 1)
	pCode("s = append(s, 1)")
	pfTree("s address(%p)", s)
	pfTree("s[0](%p)", s[0])

	slice1 = append(slice1, []int{1}...)
	pCode("slice1 = append(slice1, []int{1}...)")
	pSliceLenAndCap(slice1)
	pfTree("value: %v", slice1)
	// slice1 = append(slice1, []int{2, 3}...)
	// printfCode("slice1 = append(slice1, []int{2, 3}...)")
	// printSliceLenAndCap(slice1)
	// printfChild("value: %v", slice1)
	// slice1 = append(slice1, []int{4}...)
	// printfCode("slice1 = append(slice1, []int{4}...)")
	// printSliceLenAndCap(slice1)
	// printfChild("value: %v", slice1)
	// slice1 = append(slice1, []int{5, 6, 7}...)
	// printfCode("slice1 = append(slice1, []int{5,6,7}...)")
	// printSliceLenAndCap(slice1)
	// printfChild("value: %v", slice1)

	slice1 = make([]int, 1024)
	slice1 = append(slice1, []int{1}...)
	// printfCode("slice1 = make([]int, 1024)")
	// printfCode("slice1 = append(slice1, []int{1}...)")
	// printSliceLenAndCap(slice1)
	// modify value
	pTitle("modify()")
	slice1 = make([]int, 5)
	slice1_by_directly := slice1[:]
	slice1_by_copy := make([]int, 5)
	copy(slice1_by_copy, slice1)
	slice1[0] = 1
	// printfCode("slice1 = make([]int, 5)")
	// printfCode("slice1_by_directly := slice1[:]")
	// printfCode("slice1_by_copy := make([]int, 5)")
	// printfCode("copy(slice1_by_copy, slice1)")
	// printfCode("slice1[0] = 1")
	// printfChild("slice1[0]: %d (%p)", slice1[0], &slice1[0])
	// printfChild("slice1_by_directly[0]: %d (%p)", slice1_by_directly[0], &slice1_by_directly[0])
	// printfChild("slice1_by_copy[0]: %d (%p)", slice1_by_copy[0], &slice1_by_copy[0])
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

	println(s1)
	println(s2)
	println(slice)
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
	println("origin slice[0]: ", slice[0])

	changeFirstBeforeAppend(slice)
	println("changeFirstBeforeAppend slice[0]: ", slice[0])

	changeFirst(slice)
	println("changeFirst slice[0]: ", slice[0])
}

func TestSliceExample3(t *testing.T) {
	myAppend := func(s []int) []int {
		s = append(s, 100)
		return s
	}
	s := []int{1, 1, 1}
	newS := myAppend(s)

	println(s)
	println(newS)

	s = newS

	myAppendPtr := func(s *[]int) {
		*s = append(*s, 100)
	}
	myAppendPtr(&s)
	println(s)
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
	println("s:", s)
	println("x:", x)
	println("y:", y)
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
	println(a, x)
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
