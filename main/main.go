package main

import (
	"github.com/arvin-gao/gotutorial/child"
	"github.com/arvin-gao/gotutorial/child/export"
)

func main() {
	fExport()
	fImport()
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
