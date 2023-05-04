package gotutorial

import (
	"fmt"
	"runtime"
	"testing"
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
