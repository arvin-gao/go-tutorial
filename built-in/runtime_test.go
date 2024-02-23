package packages

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func TestMemoryStatsOnRuntime(t *testing.T) {
	var s = make([]string, 100000)
	for i := 0; i < len(s); i++ {
		s[i] = "test"
	}

	printMemoryStats()

	runtime.GC()

	printMemoryStats()
}

func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}

type Structure struct {
	UInt8    uint8 // is a byte
	Int      int
	PtrInt   *int
	Str      string
	SliceStr []string
}

// In this case, the `s := Structure{}` can not measured by this method.
func memUsage(m1, m2 *runtime.MemStats) {
	ptr("Alloc:", m2.Alloc-m1.Alloc,
		"TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc,
		"HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}

// https://dlintw.github.io/gobyexample/public/memory-and-sizeof
func TestMemory2(t *testing.T) {
	var m1, m2, m3, m4, m5, m6 runtime.MemStats

	runtime.ReadMemStats(&m1)
	var s Structure
	runtime.ReadMemStats(&m2)

	ptr("> var s Structure")
	ptr("sizeof(s.uint8)", unsafe.Sizeof(s.UInt8),
		"offset=", unsafe.Offsetof(s.UInt8))
	ptr("sizeof(s.int)", unsafe.Sizeof(s.Int),
		"offset=", unsafe.Offsetof(s.Int))
	ptr("sizeof(s.*int)", unsafe.Sizeof(s.PtrInt),
		"offset=", unsafe.Offsetof(s.PtrInt))
	ptr("sizeof(s.string)", unsafe.Sizeof(s.Str),
		"offset=", unsafe.Offsetof(s.Str))

	ptr("sizeof(s.[]string)", unsafe.Sizeof(s.SliceStr),
		"offset=", unsafe.Offsetof(s.SliceStr))

	ptr("sizeof(s)", unsafe.Sizeof(s))

	// We will see 0 bytes, because it is on stack, so sizeof is the
	// proper method to tell how much memory allocated.
	memUsage(&m1, &m2)

	// Even string assignment is in stack.
	runtime.ReadMemStats(&m3)
	str := "abc"
	runtime.ReadMemStats(&m4)
	memUsage(&m3, &m4)

	// map will alloc memory in heap
	runtime.ReadMemStats(&m5)
	t3 := map[int]string{1: "x"}
	runtime.ReadMemStats(&m6)
	memUsage(&m5, &m6)

	fmt.Println(str, t3) // prevent compiler error
}
