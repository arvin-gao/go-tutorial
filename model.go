package gotutorial

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) printMyName() {
	fmt.Println(u.Name)
}
