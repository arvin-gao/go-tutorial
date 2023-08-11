package mytesting

import (
	"strings"
	"testing"
)

type Obj struct {
	col string
}

func BenchmarkSliceAllocation(b *testing.B) {
	size := 10000000
	data := Obj{col: strings.Repeat("t", 1000)}

	var s1 []Obj
	s2 := make([]Obj, size)
	s3 := make([]Obj, 0, size)
	slices := [][]Obj{s1, s2, s3}

	putDataToEmptySliceByAppend := func(s []Obj) {
		for i := 0; i < size; i++ {
			s = append(s, data)
		}
	}

	putDataToFullCapAndLenSliceByIndex := func(s []Obj) {
		for i := 0; i < size; i++ {
			s[i] = data
		}
	}

	putDataToEmptyCapSliceByAppend := func(s []Obj) {
		for i := 0; i < size; i++ {
			s = append(s, data)
		}
	}

	cases := []struct {
		funcName string
		f        func(s []Obj)
	}{
		{"putDataToEmptySliceByAppend", putDataToEmptySliceByAppend},
		{"putDataToFullCapAndLenSliceByIndex", putDataToFullCapAndLenSliceByIndex},
		{"putDataToEmptyCapSliceByAppend", putDataToEmptyCapSliceByAppend},
	}

	for i, v := range cases {
		s := slices[i]
		b.Run(v.funcName, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				v.f(s)
			}
		})

		// release memory.
		slices[i] = nil
	}
}
