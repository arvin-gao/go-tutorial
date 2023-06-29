package testing

import "testing"

var arrLen = 10000000

func BenchmarkTakeVarLenForLoopEveryTimes(b *testing.B) {
	arr := make([]int, arrLen)
	for i := 0; i < b.N; i++ {
		count := 0
		for i := 0; i < len(arr); i++ {
			count += len(arr)
		}
	}
}

func BenchmarkTakeVarLenToVarForLoop(b *testing.B) {
	arr := make([]int, arrLen)
	for i := 0; i < b.N; i++ {
		count := 0
		arrLen := len(arr)
		for i := 0; i < arrLen; i++ {
			count += arrLen
		}
	}
}
