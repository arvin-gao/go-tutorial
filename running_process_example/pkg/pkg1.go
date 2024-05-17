package pkg

import (
	"fmt"

	"github.com/arvin-gao/gotutorial/running_process_example/srv"
)

var pkg1Variable1 = "pkg1 variable 1"

// Public variable is changeable.
var Pkg1Variable2 = "pkg1 variable 2"

func ThisIsPkg1PublicFunction() {
	fmt.Println("this is public function from pkg package of pkg1.go file")
}

func ThisIsPkg1PublicFunctionWithSrvFunction() {
	fmt.Println("this is public function with srv function from pkg package of pkg1.go file")
	srv.ThisIsSrvPublicFunction()
}

func PrintPkg1PublicVariable() {
	fmt.Println("print public variable from pkg:", Pkg1Variable2)
}

func thisIsPrivateFunction() {
	fmt.Println("this is private function from pkg package of pkg1.go file")
}

func init() {
	fmt.Println("run init function from pkg/pkg1.go")
}
