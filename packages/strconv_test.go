package packages

import (
	"strconv"
	"testing"
)

// refer to: https://gobyexample.com/number-parsing
func TestStrconv(t *testing.T) {
	f, _ := strconv.ParseFloat("1.234", 64)
	ptr(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	ptr(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	ptr(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	ptr(u)

	k, _ := strconv.Atoi("135")
	ptr(k)

	_, e := strconv.Atoi("wat")
	ptr(e)
}
