package gotutorial

import (
	"testing"
)

func TestMap(t *testing.T) {
	// Declaring empty map
	var _ = map[string]int{}
	// Initializing a map
	var _ = map[string]int{"A": 1, "B": 2}
	// With make()
	m1 := make(map[string]int)
	m2 := make(map[int]string, 2)
	m3 := make(map[any]User)

	var m11 map[string]int
	m22 := map[string]int{"k1": 1, "k2": 2, "k3": 3}
	pass(m11, m22)

	m1["K1"] = 7
	ptr(m1["k1"])

	// map length.
	ptr("m2 length:", len(m2))
	m2[1] = "a"
	ptr("appended m2 length:", len(m2))

	m3[true] = User{}

	// Check the key is exist.
	val := m22["k2"]
	val, isExist := m22["k2"]
	if isExist {
		ptr("has key, value:", val) // 2
	} else {
		ptr("non key, value:", val) // 0
	}

	// Delete item by the key.
	delete(m1, "k1")
	ptr("k1 value:", m1["k1"])
}

/*
when using the make function to create a map,
the second argument is neither the initial length nor the capacity of the result map.

It is just a hint for Go runtime to allocate a large enough backing array to hold at least the specified number of entries.

The length (number of entries) of the result map is zero.
*/
func TestMap2(t *testing.T) {
	m := make(map[string]int, 3)
	len1 := len(m)

	m["k1"] = m["k2"]

	len2 := len(m)

	ptr(len1, len2)
}

func TestChangeMapValue(t *testing.T) {
	m1 := make(map[string]User)
	m1["user1"] = User{Age: 1}
	user1 := m1["user1"]
	user2 := m1["user1"]
	pPtr(&user1)
	pPtr(&user2)

	m2 := make(map[string]*User)
	m2["user1"] = &User{Age: 1}
	m2["user1"].Age = 5

	// m1["user1"].Age = 23 // panic, cannot assign to struct field m1["user1"].Age in map
	user1.Age = 12
	m1["user1"] = user1

	ptrlnf("%+v\n%+v", m1["user1"], m2["user1"])
}
