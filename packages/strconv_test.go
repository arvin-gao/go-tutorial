package packages

import (
	"strconv"
	"testing"
)

// refer to: https://gobyexample.com/number-parsing
func TestStrconvStringToType(t *testing.T) {
	// to float64.
	strconv.ParseFloat("1.234", 64)
	// to int64.
	strconv.ParseInt("123", 0, 64)
	// to int and with error.
	strconv.Atoi("135")
	// Hexadecimal/Octal to int64.
	strconv.ParseInt("0x1c8", 0, 64)
	// to uint64.
	strconv.ParseUint("789", 0, 64)
	// to Boolean
	strconv.ParseBool("true")
}

func TestStrconvTypesToString(t *testing.T) {
	// Float to String
	var flo float64 = 3.1415926535
	strconv.FormatFloat(flo, 'E', -1, 32)
}
