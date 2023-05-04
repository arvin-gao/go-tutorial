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
		fmt.Println("1", a)
	}
	fmt.Println(a)
}
