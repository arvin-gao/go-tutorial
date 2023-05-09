package packages

import (
	"strconv"
	"testing"
)

// refer to: https://gobyexample.com/number-parsing
func TestStrconv(t *testing.T) {
	f, _ := strconv.ParseFloat("1.234", 64)
	pln(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	pln(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	pln(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	pln(u)

	k, _ := strconv.Atoi("135")
	pln(k)

	_, e := strconv.Atoi("wat")
	pln(e)
}
