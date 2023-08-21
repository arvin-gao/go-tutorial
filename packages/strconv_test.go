package packages

import (
	"strconv"
	"testing"
)

// refer to: https://gobyexample.com/number-parsing
func TestStrconv(t *testing.T) {
	// to float64.
	strconv.ParseFloat("1.234", 64)
	// to int64.
	strconv.ParseInt("123", 0, 64)
	// Hexadecimal/Octal to int64.
	strconv.ParseInt("0x1c8", 0, 64)
	// to uint64.
	strconv.ParseUint("789", 0, 64)
	// to int.
	strconv.Atoi("135")
	// got error.
	strconv.Atoi("wat")
}
