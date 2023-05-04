package gotutorial

import (
	"fmt"
	"testing"
)

// TODO: map len and capacity.
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

/*
when using the make function to create a map, the second argument is neither the initial length nor the capacity of the result map.
It is just a hint for Go runtime to allocate a large enough backing array to hold at least the specified number of entries.

The length (number of entries) of the result map is zero. m["Go"] = m["Go"] is equivalent to m["Go"] = 0.
After the assignment, the map contains one entry.
*/
func TestMap2(t *testing.T) {
	m := make(map[string]int, 3)
	len1 := len(m)

	m["Go"] = m["Go"]

	len2 := len(m)

	println(len1, len2)
}
