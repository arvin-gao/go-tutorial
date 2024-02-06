package gotutorial

import (
	"strconv"
	"testing"
)

func TestConvertType(t *testing.T) {
	// numbers
	var num int = 1
	numA := int64(num)
	numB := float64(num)
	_ = numA + int64(numB)

	// string.
	bys := []byte("abc")
	str := string(bys)
	char := str[0] // byte [0]byte
	str = strconv.Itoa(123)
	_ = str + str + string(char)

	// array & slice
	var arr [1]int
	slice := arr[:]

	arr2 := [1]int(slice)
	_, _ = arr[0], slice[0]

	// interface | any #可接受任何值
	var i any
	i = bys
	_bys, _ := i.([]byte)
	i = arr2
	_arr2 := i.([1]int)

	pass(_bys, _arr2)
}
