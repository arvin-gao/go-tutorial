package gotutorial

import (
	"testing"
)

func TestIfElse(t *testing.T) {
	if 7%2 == 0 {
		ptr("7%2 == 0")
	} else {
		ptr("7%2 != 0")
	}

	var ok bool = true
	if ok {
		ptr("ok")
	}

	num := 9
	if num < 0 {
		ptr("num < 0")
	} else if num < 10 {
		ptr("num < 10")
	} else {
		ptr("num > 10")
	}
}

func TestConditionVariable(t *testing.T) {
	var a = 1
	ptr(a) // 1

	// 1. Short variable declarations
	if a := 2; a == 2 {
		// inside
		ptr(a) // 2
		a = 3
		ptr(a) // 3
	}
	// outside
	ptr(a) // 1

	// 2.
	if a = 4; a == 4 {
		ptr(a) // 4
	}
	ptr(a) // 4

	// 3.
	if a = 5; a == 5 {
		a = 6
		ptr(a) // 6
	}
	ptr(a) // 6
}

func TestSwitch(t *testing.T) {
	// variable
	var num = 11
	switch num {
	case 9:
		ptr("num == 9")
		// Multiple cases
	case 10, 11:
		ptr("num == 10 or 11")
	default:
		ptr("num != 9 && num != 10")
	}

	switch num {
	case 10, 11:
		ptr("num == 10 or 11")
		// The fallthrough statement will forces the execution of the next statement.
		fallthrough
	case 12:
		ptr("12")
	case 13:
		ptr("13")
	default:
		ptr("num != 9 && num != 10")
	}

	// Using conditional statements.
	num = 10
	switch {
	case num > 0 && num < 9:
		ptr("num > 0 && num < 9")
	case num > 9:
		ptr("num > 9")
	case num == 10:
		ptr("num == 10")
	}
	// Initializer statement
	switch number := 5; {
	case number < 5:
		ptr("Smaller than 5")
	case number == 5:
		ptr("Five")
	}

	// With enumeration variables.
	type myEnum int
	var (
		e1 myEnum
		e2 myEnum
	)
	var req = e1
	switch req {
	case e1:
		ptr("req == enumeration 1")
	case e2:
		ptr("req == enumeration 2")
	}

	// Check the type of the variable.
	var v any = num
	switch t := v.(type) {
	case int:
		ptr("v is int")
	case int64:
		ptr("v is int64")
	case string:
		ptr("v is string")
	default:
		ptr("The v type is:", t)
	}
}
