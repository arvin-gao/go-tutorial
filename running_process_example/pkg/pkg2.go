package pkg

import "fmt"

const pkg2Const1 = "pkg2Const1"
const Pkg2Const2 = "pkg2Const2"

func ThisIsPkgPublicFunction2() {
	fmt.Println("this is public function from pkg package of pkg2.go file")
}

func init() {
	fmt.Println("run init function from pkg/pkg2.go")
}
