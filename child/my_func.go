package child

func MyFunc() {
	println("call MyFunc")
}

func init() {
	GlobalVariable++
	println("my_func init. GlobalVariable:", GlobalVariable)
}

func init() {
	println("init 2")
}
