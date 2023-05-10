package gotutorial

import (
	"errors"
	"strconv"
	"testing"
)

func TestDeclaration(t *testing.T) {
	var a int
	var b int = 1
	var c = 1
	var d *int
	e := 1

	var a1, b1, c1 = 1, 2.0, "3"

	var (
		a2 int
		b2 int
	)

	v1, v2, err := getMultiVars()
	_, v22, _ := getMultiVars()

	const cInt = 1
	const cStr string = "1"

	pass(a2, b2, a, b, c, d, cInt, cStr, v1, v2, v22, err, a1, b1, c1, e)
}

func TestInitNil(t *testing.T) {
	var (
		ch       chan string
		m        map[string]interface{}
		in       interface{}
		num      *int
		user     *User
		slice    []int
		ptrSlice *[]int
	)
	println(
		ch == nil,
		m == nil,
		in == nil,
		num == nil,
		user == nil,
		slice == nil,
		ptrSlice == nil,
	)
}

func TestConvertType(t *testing.T) {
	// numbers
	var num = 1
	numA := int64(num)
	numB := float64(num)
	_ = numA + int64(numB)

	// string.
	str := string([]byte("abc"))
	bys := []byte("abc")
	char := str[0] // byte
	str = strconv.Itoa(123)
	_ = str + str + string(char)

	// array & slice
	var arr [1]int
	slice := arr[:]
	arr2 := [1]int(slice)
	_, _ = arr[0], slice[0]

	// interface | any
	var i any
	i = bys
	_bys, _ := i.([]byte)
	i = arr2
	_arr2 := i.([1]int)

	pass(_bys, _arr2)
}

func getMultiVars() (int, string, error) {
	return 0, "", errors.New("err")
}
