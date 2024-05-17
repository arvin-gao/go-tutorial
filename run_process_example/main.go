package main

import (
	"fmt"

	"github.com/arvin-gao/gotutorial/running_process_example/pkg"
	"github.com/arvin-gao/gotutorial/running_process_example/utils"
)

func main() {
	fmt.Println("run main function from main.go")
	pkg.ThisIsPkg1PublicFunction()
	pkg.ThisIsPkgPublicFunction2()
	utils.ThisIsUtilsPublicFunction()
	fmt.Println("print public variable from pkg", pkg.Pkg1Variable2)
	fmt.Println("print public const from pkg", pkg.Pkg2Const2)
}

func init() {
	fmt.Println("run init function from main.go")
}
