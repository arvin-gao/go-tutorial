package testing

import (
	"fmt"
	"testing"
)

func BenchmarkConcatStringBetweenNormalAndSprintf(b *testing.B) {
	concatStringByPlus := func(s1, s2 string) string {
		return s1 + "." + s2 + "."
	}

	concatStringBySprintf := func(s1, s2 string) string {
		return fmt.Sprintf("%s.%s.", s1, s2)
	}

	cases := []struct {
		funcName string
		f        func(s1, s2 string) string
	}{
		{"plus", concatStringByPlus},
		{"sprintf", concatStringBySprintf},
	}

	for _, v := range cases {
		b.Run(v.funcName, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				v.f("s1", "s2")
			}
		})
	}
}
