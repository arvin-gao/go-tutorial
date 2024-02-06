package gotutorial

import (
	"bytes"
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// Initialize array.
	var arr [2]int                           // [0,0], len:2 , cap:2
	var _ = [2]int{}                         // [0,0], len:2 , cap:2
	var _ = [2]int{1, 2}                     // [1,2], len:2 , cap:2
	var _ = [...]int{1, 2, 2, 3, 3, 4, 4, 5} // [1,2], len:8 , cap:8
	_ = [3][3]int{                           // [[1 2 3] [2 3 4] [3 4 5]], len: 3, cap: 3
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	// Initialize values for specific array elements
	_ = [5]int{1: 1, 3: 25} // [0 1 0 25 0], len:5 , cap:5

	ptr("val:", arr)
	ptr("len:", len(arr))
	ptr("cap:", cap(arr))

	// Copy array
	x := [5]int{0, 5, 10, 15, 20}
	// Copy array values
	y := x
	// Copy by reference
	z := &x
	x[0] = 1
	ptrlnf("x:%p", &x[0])
	ptrlnf("y:%p", &y[0])
	ptrlnf("z:%p", &z[0])
}

func TestSlice(t *testing.T) {
	// Initialize slice.
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
	ptrSubject("var ss = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}")
	ptrlnf("ss[:] -> %v", ss[:])
	ptrlnf("ss[4:] -> %v", ss[4:])
	ptrlnf("ss[:3] -> %v", ss[:3])
	ptrlnf("ss[2:4] -> %v", ss[2:4])

	var ss2 []int = nil
	ptrSubject("var ss2 []int = nil")
	ptrlnf("ss2 address is %p", ss2)
	ss2 = append(ss2, 1)
	ptrSubject("ss2 = append(ss2, 1)")
	ptrlnf("ss2 address is %p", ss2)
	// len cap
	pTitle("init slice")
	ptrSubject("var slice1 []int")
	ptr("val:", slice1)
	ptrSliceLenAndCap(slice1)
	ptrSubject("var slice2 = []int{}")
	ptr("val:", slice2)
	ptrSliceLenAndCap(slice2)
	ptrSubject("var slice3 = []int{1, 2}")
	ptr("val:", slice3)
	ptrSliceLenAndCap(slice3)
	ptrlnf("slice3[0]: %p", &slice3[0])
	ptrSubject("var slice4 = make([]int, 2)")
	ptr("val:", slice4)
	ptrSliceLenAndCap(slice4)
	ptrSubject("var slice5 = make([]int, 2, 3)")
	ptr("val:", slice5)
	ptrSliceLenAndCap(slice5)
	ptrSubject("var slice6 = make([]int, 0, 3)")
	ptr("val:", slice6)
	ptrSliceLenAndCap(slice6)
	ptrSubject("var slice7 = make([]int, len(slice3))")
	ptrfTree("copy(slice7, slice3)")
	ptr("val:", slice7)
	ptrSliceLenAndCap(slice7)
	ptrlnf("slice7[0]: %p", &slice7[0])
	ptrSubject("slice8 := slice7[:len(slice7)-1]")
	ptr("val:", slice8)
	ptrSliceLenAndCap(slice8)
	ptrlnf("slice8[0]: %p", &slice8[0])

	pTitle("append()")
	// TODO: 檢查 append 後的slice地址是否相同？.
	s := make([]int, 2, 3)
	ptrSubject("s := make([]int, 2, 3)")
	ptrlnf("s[0]: %p", &s[0])
	s = append(s, 1)
	ptrSubject("s = append(s, 1)")
	ptrlnf("s[0]: %p", &s[0])
	s = append(s, 1)
	ptrSubject("s = append(s, 1)")
	ptrlnf("s[0]: %p", &s[0])

	ptrSubject("slice1 = append([]int{}, []int{1,2,3}...)")
	slice1 = append([]int{}, []int{1, 2, 3}...)
	ptr("val:", slice1)
	ptrSliceLenAndCap(slice1)
	slice1 = append(slice1, []int{2, 3}...)
	ptrSubject("slice1 = append(slice1, []int{2, 3}...)")
	ptr("val:", slice1)
	ptrSliceLenAndCap(slice1)
	slice1 = append(slice1, []int{4}...)
	ptrSubject("slice1 = append(slice1, []int{4}...)")
	ptrSliceLenAndCap(slice1)
	ptr("val:", slice1)
	slice1 = append(slice1, []int{5, 6, 7}...)
	ptrSubject("slice1 = append(slice1, []int{5,6,7}...)")
	ptr("val:", slice1)
	ptrSliceLenAndCap(slice1)

	slice1 = make([]int, 1024)
	slice1 = append(slice1, 1)
	ptrSubject("slice1 = make([]int, 1024)")
	ptrfTree("slice1 = append(slice1, 1)")
	ptrSliceLenAndCap(slice1)
	// modify value
	pTitle("modify()")
	slice1 = make([]int, 5)
	slice1BySlicingAll := slice1[:]
	slice1ByCopy := make([]int, 5)
	copy(slice1ByCopy, slice1)
	slice1[0] = 1
	ptrSubject("slice1 = make([]int, 5)")
	ptrfTree("slice1BySlicingAll := slice1[:]")
	ptrfTree("slice1ByCopy := make([]int, 5)")
	ptrfTree("copy(slice1ByCopy, slice1)")
	ptrfTree("slice1[0] = 1")
	ptrlnf("slice1[0]: %d, addr[0]: %p", slice1[0], &slice1[0])
	ptrlnf("slice1BySlicingAll[0]: %d, addr[0]: %p", slice1BySlicingAll[0], &slice1BySlicingAll[0])
	ptrlnf("slice1ByCopy[0]: %d, addr[0]: %p", slice1ByCopy[0], &slice1ByCopy[0])
	slice1BySlicingAll = append(slice1BySlicingAll, 1)
	slice1[0] = 2
	ptrSubject("slice1BySlicingAll = append(slice1BySlicingAll, 1)")
	ptrfTree("slice1[0] = 2")
	ptrlnf("slice1[0]: %d, addr[0]: %p", slice1[0], &slice1[0])
	ptrlnf("slice1BySlicingAll[0]: %d, addr[0]: %p", slice1BySlicingAll[0], &slice1BySlicingAll[0])

	changeFirstValue := func(s []int) {
		s[0] = 100
	}
	slice1[0] = 1
	ptrSubject("slice1[0] = 1")
	ptrlnf("slice1[0]: %d, addr: %p", slice1[0], &slice1[0])
	changeFirstValue(slice1)
	ptrSubject("changeFirstValue := func(s []int) { s[0] = 100 }")
	ptrfTree("changeFirstValue(slice1)")
	ptrlnf("slice1[0]: %d, addr: %p", slice1[0], &slice1[0])
}

func RangeSlice() {
	nums := []int{1, 2, 3}
	// Basic loop.
	for i := 0; i < len(nums); i++ {
		// nums[i]
	}

	// Loop-of
	for idx, num := range nums {
		ptr(idx, num)
	}

	// Loop with increase index
	i := 0
	for range nums {
		ptr(nums[i])
		i++
	}
}

func TestSliceAdvance(t *testing.T) {
	// This behavior can lead to unexpected memory usage.
	data := func() []int {
		raw := make([]int, 10000)
		ptr("raw length:", len(raw))
		ptr("raw capacity:", cap(raw))
		ptrlnf("raw[0] address: %p", &raw[0])
		return raw[:3]
		// Solution.
		// res := make([]int, 3)
		// copy(res, raw[:3])
		// return res
	}()
	ptr("data length:", len(data))   // 3
	ptr("data capacity:", cap(data)) // 10000
	ptrlnf("data[0] address: %p", &data[0])
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
	ptrfTree("x = %v", x)

	y := append(s, 22)
	pCode("y := append(s, 22)")
	ptrfTree("y = %v", y)

	z := []int{2: 5, 6, 0: 7}
	pCode("z := []int{2: 5, 6, 0: 7}")
	ptrfTree("z = %v", z)

	y[len(y)-1] = 23
	pCode("y[len(y)-1] = 23")
	ptrfTree("s = %v", s)
	ptrfTree("x = %v", x)
	ptrfTree("y = %v", y)
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

func TestDirPathProblem(t *testing.T) {
	pTitle("Both slices referenced the same underlying array data from the original slice")
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	ptrSubject("path := []byte(\"AAAA/BBBBBBBBB\")")
	ptrfTree("sepIndex := bytes.IndexByte(path, '/')")
	ptrfTree("val: %s", string(path))
	ptrfTree("len: %d", len(path))
	ptrfTree("cap: %d", cap(path))
	ptr("dir1 := path[:sepIndex]")
	ptrfTree("val: %s", string(dir1))
	ptrfTree("len: %d", len(dir1))
	ptrfTree("cap: %d", cap(dir1))
	ptr("dir2 := path[sepIndex+1:]")
	ptrfTree("val: %s", string(dir2))
	ptrfTree("len: %d", len(dir2))
	ptrfTree("cap: %d", cap(dir2))

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	ptr("dir1")
	ptrfTree("val:", string(dir1))
	ptrfTree("len:", len(dir1))
	ptrfTree("cap:", cap(dir1))
	ptr("dir2")
	ptrfTree("val:", string(dir2))
	ptrfTree("len:", len(dir2))
	ptrfTree("cap:", cap(dir2))
	ptr("new path =>", string(path))
}

func TestOldSliceProblem(t *testing.T) {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // 3 3 [1 2 3]
	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // 2 2 [2 3]
	for i := range s2 {
		s2[i] += 20
	}
	fmt.Println(s1) // [1 22 23]
	fmt.Println(s2) // [22 23]

	s2 = append(s2, 4)
	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println(s1) // [1 22 23]
	fmt.Println(s2) // [32 33 14]
}
