package packages

import "fmt"

var (
	p    = fmt.Print
	ptrf = fmt.Printf
	ptr  = fmt.Println
)

func pass(v ...any) {}
