package gotutorial

import (
	"testing"
)

func TestIfElse(t *testing.T) {
	if 7%2 == 0 {
		print("7%2 == 0")
	} else {
		print("7%2 != 0")
	}

	var ok bool = true
	if ok {
		print("ok")
	}

	if num := 9; num < 0 {
		print("num < 0")
	} else if num < 10 {
		print("num < 10")
	} else {
		print("num > 10")
	}
	var num2 int
	if num2 = 9; num2 < 0 {
	}
}

func TestSwitch(t *testing.T) {
	// variable
	var num = 11
	switch num {
	case 9:
		println("num == 9")
	case 10, 11:
		println("num == 10 or 11")
	default:
		println("num != 9 && num != 10")
	}

	switch num {
	case 10, 11:
		println("num == 10 or 11")
		fallthrough
	case 12:
		println("12")
	case 13:
		println("13")
	default:
		println("num != 9 && num != 10")
	}

	// any or interface type
	var v any = num
	switch myType := v.(type) {
	case int:
		println("v is int")
	case int64:
		println("v is int64")
	case string:
		println("v is string")
	default:
		println("my type is:", myType)
	}

	num = 10
	switch {
	case num > 0 && num < 9:
		println("num > 0 && num < 9")
	case num > 9:
		println("num > 9")
	case num == 10:
		println("num == 10")
	}

	// enum
	type myEnum int
	var (
		enum1 myEnum
		enum2 myEnum
	)
	var req = enum1
	switch req {
	case enum1:
		println("req == enum1")
	case enum2:
		println("req == enum2")
	}
}
