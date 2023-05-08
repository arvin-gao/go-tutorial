package gotutorial

import (
	"testing"

	"github.com/arvin-gao/gotutorial/child"
	"github.com/arvin-gao/gotutorial/child/export"
	"github.com/arvin-gao/test2"
	. "github.com/arvin-gao/test2"
	test22 "github.com/arvin-gao/test2"
)

func TestPublicAndPrivate(t *testing.T) {
	export.Tool1()
}

// find package -> var -> init() -> main()
func TestInitPackage(t *testing.T) {
	child.MyFunc()
	println("child.GlobalVariable:", child.GlobalVariable)
	child.MyFunc()
}

func TestImport(t *testing.T) {
	test2.RunTest2()
	RunTest2()
	test22.RunTest2()
}
