package main

import (
	"fmt"
)

func myFunc(myDoneArg *bool) {
	*myDoneArg = true
}
func main() {

	myDone := false

	go myFunc(&myDone)

	for !myDone {
		// runtime.Gosched()
		fmt.Println("ptr loop")
	}
	fmt.Println("ptr outside")
}
