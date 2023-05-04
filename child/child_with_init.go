package child

import "fmt"

func Child() {
	fmt.Println("child func")
}

func init() {
	GlobalVar = 2
	fmt.Println("child with init. GlobalVar: ", GlobalVar)
}
