package gotutorial

import (
	"sync"
	"testing"
)

// TODO: map len and capacity.
func TestMap(t *testing.T) {
	m1 := make(map[string]int)
	// TODO:
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
	ptr("m2 length:", len(m2))

	m3[true] = User{}

	// check exist.
	val := m22["k2"]
	val, isExist := m22["k2"]
	if isExist {
		ptr("has key, value:", val) // 2
	} else {
		ptr("non key, value:", val) // 0
	}

	// delete the key.
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

	ptrf("%+v\n%+v", m1["user1"], m2["user1"])
}

// TODO: advance.
func TestConcurrencyMap(t *testing.T) {
	var wg sync.WaitGroup

	builtinMap := make(map[string]int)

	wg.Add(2)
	// issue: fatal error: concurrent map writes
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			builtinMap["key"]++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 2000; i++ {
			builtinMap["key"]++
		}
	}()

	wg.Wait()
	ptr(builtinMap["key"])
}

// TODO:
func TestSyncMap(t *testing.T) {
	var wg sync.WaitGroup

	// ?
	var safeMap sync.Map
	safeMap.Store("k1", "v1")
	safeMap.Store("k2", "v2")
	v, ok := safeMap.Load("k1")
	ptr(v, ok)
	return
	wg.Add(2)
	// issue: fatal error: concurrent map writes
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			safeMap.Store("k1", "v1")
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 2000; i++ {
			// safeMap.Swap(key any, value any)
		}
	}()

	wg.Wait()

}
