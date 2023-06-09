package packages

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	pln("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	pln("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	pln("Sorted:", s)
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func TestCustomSort(t *testing.T) {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	pln(fruits)
}
