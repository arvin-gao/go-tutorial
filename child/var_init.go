package child

import "github.com/arvin-gao/gotutorial/child/child2"

var GlobalVariable int = 10

func init() {
	GlobalVariable++
	println("var init. GlobalVariable:", GlobalVariable)
	child2.Call()
}
