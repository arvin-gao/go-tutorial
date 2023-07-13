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
		ptr("num == 9")
	case 10, 11:
		ptr("num == 10 or 11")
	default:
		ptr("num != 9 && num != 10")
	}

	switch num {
	case 10, 11:
		ptr("num == 10 or 11")
		fallthrough
	case 12:
		ptr("12")
	case 13:
		ptr("13")
	default:
		ptr("num != 9 && num != 10")
	}

	// any or interface type
	var v any = num
	switch myType := v.(type) {
	case int:
		ptr("v is int")
	case int64:
		ptr("v is int64")
	case string:
		ptr("v is string")
	default:
		ptr("my type is:", myType)
	}

	num = 10
	switch {
	case num > 0 && num < 9:
		ptr("num > 0 && num < 9")
	case num > 9:
		ptr("num > 9")
	case num == 10:
		ptr("num == 10")
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
		ptr("req == enum1")
	case enum2:
		ptr("req == enum2")
	}
}

func TestConditionVariable(t *testing.T) {
	var a = 1
	ptr("a =", a)
	if a := 2; a == 2 {
		ptr(a)
		a = 3
		ptr(a)
	}
	ptr("a =", a)

	if a = 4; a == 4 {
		pass()
	}
	ptr("a =", a)

	if a = 5; a == 5 {
		a = 6
	}
	ptr("a =", a)
}
