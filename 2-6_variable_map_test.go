package gotutorial

import (
	"sync"
	"testing"
)

// TODO: map len and capacity.
func TestMap(t *testing.T) {
	m1 := make(map[string]int)
	m2 := make(map[int]string, 2)
	m3 := make(map[any]User)

	var m11 map[string]int
	m22 := map[string]int{"k1": 1, "k2": 2, "k3": 3}
	pass(m11, m22)

	m1["k1"] = 7
	println(m1["k1"])

	// map length.
	println("m2 length:", len(m2))
	m2[1] = "a"
	println("m2 length:", len(m2))

	m3[true] = User{}

	// check exist.
	v, ok := m1["k2"]
	if ok {
		println("has key, value:", v)
	} else {
		println("non key, value:", v)
	}

	// delete the key.
	delete(m1, "k1")
	println("k1 value:", m1["k1"])
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

	println(len1, len2)
}

func TestChangeMapValue(t *testing.T) {
	m1 := make(map[string]User)
	m1["user1"] = User{Age: 1}
	user11 := m1["user1"]
	user12 := m1["user1"]
	pPtr(&user11)
	pPtr(&user12)

	m2 := make(map[string]*User)
	m2["user1"] = &User{Age: 1}
	m2["user1"].Age = 5

	println(m2["user1"].Age)

}

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
	println(builtinMap["key"])
}

func TestSyncMap(t *testing.T) {
	var wg sync.WaitGroup

	// ?
	var safeMap sync.Map
	safeMap.Store("k1", "v1")
	safeMap.Store("k2", "v2")
	v, ok := safeMap.Load("k1")
	println(v, ok)
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
