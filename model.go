package gotutorial

type User struct {
	Name string
	Age  int
}

func (u User) printMyName() {
	println(u.Name)
}
