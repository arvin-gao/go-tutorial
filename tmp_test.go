package gotutorial

import (
	"fmt"
	"testing"
)

func TestTemp(t *testing.T) {
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", &p) // prints address
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the pointer p: %p\n", &s) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

	var a int
	if a = 10; a > 9 {
		ptr("1", a)
	}
	ptr(a)
}

type Obj struct {
	Name string
}

func (a *Obj) printObjWithPointer() {
	fmt.Printf("*a: %p\n", a)
}

func (a Obj) printObjWithoutPointer() {
	fmt.Printf("a: %p\n", &a)
}

func TestTemp2(t *testing.T) {
	var v Obj
	fmt.Printf("v: %p\n", &v)
	v.printObjWithPointer()
	ptr()
	v.printObjWithoutPointer()
	v.printObjWithoutPointer()
}

func TestTemp3(t *testing.T) {
	var values = []int{1, 2, 3, 4, 5}
	for ix := range values {
		go func(val int) {
			fmt.Print(val, " ")
		}(values[ix])
	}
}

func TestPf(t *testing.T) {
	ptrf("as%d", 12)
}
