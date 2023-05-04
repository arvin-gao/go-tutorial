package gotutorial

import (
	"fmt"
	"testing"

	"github.com/arvin-gao/gotutorial/child"
	"github.com/arvin-gao/test2"
	. "github.com/arvin-gao/test2"
	test22 "github.com/arvin-gao/test2"
)

func TestImport(t *testing.T) {
	test2.RunTest2()
	RunTest2()
	test22.RunTest2()
}

func TestInitPackage(t *testing.T) {
	child.Child()
	fmt.Println(child.GlobalVar)
	child.Child()
}
