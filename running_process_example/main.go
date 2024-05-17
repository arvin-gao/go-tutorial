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
	pkg.ThisIsPkg1PublicFunctionWithSrvFunction()
	utils.ThisIsUtilsPublicFunction()
	fmt.Println("print public const from pkg", pkg.Pkg2Const2)
	fmt.Println("print public variable of pkg package from main function:", pkg.Pkg1Variable2)
	pkg.Pkg1Variable2 = "Change pkg public variable"
	fmt.Println("print public variable of pkg package from main function:", pkg.Pkg1Variable2)
	pkg.PrintPkg1PublicVariable()
}

func init() {
	fmt.Println("run init function from main.go")
}
