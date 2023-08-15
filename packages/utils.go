package packages

import "fmt"

var (
	p   = fmt.Print
	pf  = fmt.Printf
	ptr = fmt.Println
)

func pass(v ...any) {}
