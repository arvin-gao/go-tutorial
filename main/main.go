package main

import "github.com/arvin-gao/gotutorial/child"

func main() {
	child.MyFunc()
	println("child.GlobalVariable:", child.GlobalVariable)
	child.MyFunc()
}
