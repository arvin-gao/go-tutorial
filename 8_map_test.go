package gotutorial

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var (
		a int
		b int
	)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	t2(a, b)
}

func t2(s ...int) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%p\n", &s[i])
	}
}
