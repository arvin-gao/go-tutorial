package gotutorial

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		inA, inB, want int
	}{
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, c := range cases {
		got := Sum(c.inA, c.inB)
		if got != c.want {
			t.Errorf("Sum(%d, %d), want %d, got %d", c.inA, c.inB, c.want, got)
		}
	}
}
