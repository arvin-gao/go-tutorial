package main

import (
	"github.com/arvin-gao/gotutorial/child"
	"github.com/arvin-gao/gotutorial/child/export"
	"github.com/arvin-gao/test2"

	. "github.com/arvin-gao/test2"
	test22 "github.com/arvin-gao/test2"
)

func main() {
	fExport()
	fImport()
	fPackage()
}

func fExport() {
	export.Tool1()
}

// find package -> var -> init() -> main()
func fImport() {
	child.MyFunc()
	println("child.GlobalVariable:", child.GlobalVariable)
	child.MyFunc()
}

func fPackage() {
	test2.RunTest2()
	RunTest2()
	test22.RunTest2()
}
