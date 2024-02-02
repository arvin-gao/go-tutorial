package gotutorial

import (
	"fmt"
	"testing"
)

/*
An interface(any) value with a nil non-interface dynamic value is not a nil interface.
So the function always return a non-nil interface value.

A call to the new builtin function always returns a pointer pointing to a zero value.
That means dereferencing the pointer results in a zero value.
The zero value of the *int type is nil.
*/
func TestFuncNil(t *testing.T) {
	var a = *new(*int)
	var b *int = nil
	var c = func() any {
		return b
	}()
	var d any = b
	// var *
	fmt.Printf("a:%v, b:%v, c:%v, d:%v\n", a, b, c, d)
	fmt.Printf("Is nil? -> a:%v, b:%v, c:%v, d:%v\n", a == nil, b == nil, c == nil, d == nil)
}

/*
The nil in f(nil...) is treated as a nil slice (of type []interface{}).
The nil in f(nil) is treated as an element in a slice. That is equivalent to f([]interface{}{nil}...).
The call f() is equivalent to f(nil...).
*/
func TestInterfaceSliceOnFunctionParameter(t *testing.T) {
	f := func(v ...interface{}) {
		fmt.Println(len(v))
	}
	f()
	f(nil...)
	f(nil)
	f(nil, nil)
}
