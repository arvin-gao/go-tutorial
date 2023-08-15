package gotutorial

import "testing"

func TestScopes(t *testing.T) {
	var a, b, c int = 1, 1, 1

	ptrf("a:%d,b:%d,c:%d", a, b, c)

	{
		a := 2
		b = 2
		if c = 2; c == 2 {
			pass()
		}
		var d int = 2
		ptrf("block - a:%d, b:%d, c:%d, d:%d", a, b, c, d)
	}

	// println(d) // the d is undefined.
	ptrf("a:%d,b:%d,c:%d", a, b, c)
}
