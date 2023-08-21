package packages

import (
	"strings"
	"testing"
)

// refer to: https://gobyexample.com/string-functions
func TestStrings(t *testing.T) {
	ptr(
		// true
		strings.Contains("test", "es"),
		// 2
		strings.Count("test", "t"),
		// true
		strings.HasPrefix("test", "te"),
		// true
		strings.HasSuffix("test", "st"),
		// 1
		strings.Index("test", "e"),
		// a-b
		strings.Join([]string{"a", "b"}, "-"),
		// aaaaa
		strings.Repeat("a", 5),
		// f00. replace all.
		strings.Replace("foo", "o", "0", -1),
		// f0o. replace one character.
		strings.Replace("foo", "o", "0", 1),
		// [a b c d e]
		strings.Split("a-b-c-d-e", "-"),
		// test
		strings.ToLower("TESt"),
		// TEST
		strings.ToUpper("tesT"),
	)
}
