package packages

import (
	"strings"
	"testing"
)

// refer to: https://gobyexample.com/string-functions
func TestStrings(t *testing.T) {

	pln("Contains:  ", strings.Contains("test", "es"))        // Contains:   true
	pln("Count:     ", strings.Count("test", "t"))            // Count:      2
	pln("HasPrefix: ", strings.HasPrefix("test", "te"))       // HasPrefix:  true
	pln("HasSuffix: ", strings.HasSuffix("test", "st"))       // HasSuffix:  true
	pln("Index:     ", strings.Index("test", "e"))            // Index:      1
	pln("Join:      ", strings.Join([]string{"a", "b"}, "-")) // Join:      a-b
	pln("Repeat:    ", strings.Repeat("a", 5))                // Repeat:      aaaaa
	pln("Replace:   ", strings.Replace("foo", "o", "0", -1))  // Replace:    f00
	pln("Replace:   ", strings.Replace("foo", "o", "0", 1))   // Replace:    f0o
	pln("Split:     ", strings.Split("a-b-c-d-e", "-"))       // Split:      [a b c d e]
	pln("ToLower:   ", strings.ToLower("TEST"))               // ToLower:    test
	pln("ToUpper:   ", strings.ToUpper("test"))               // ToUpper:    TEST
}
