package child

import "fmt"

var GlobalVar int = 0

func init() {
	GlobalVar = 1
	fmt.Println("another init. GlobalVar: ", GlobalVar)
}
