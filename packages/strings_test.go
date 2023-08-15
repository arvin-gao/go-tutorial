package packages

import (
	"strings"
	"testing"
)

// refer to: https://gobyexample.com/string-functions
func TestStrings(t *testing.T) {

	ptr("Contains:  ", strings.Contains("test", "es"))        // Contains:   true
	ptr("Count:     ", strings.Count("test", "t"))            // Count:      2
	ptr("HasPrefix: ", strings.HasPrefix("test", "te"))       // HasPrefix:  true
	ptr("HasSuffix: ", strings.HasSuffix("test", "st"))       // HasSuffix:  true
	ptr("Index:     ", strings.Index("test", "e"))            // Index:      1
	ptr("Join:      ", strings.Join([]string{"a", "b"}, "-")) // Join:      a-b
	ptr("Repeat:    ", strings.Repeat("a", 5))                // Repeat:      aaaaa
	ptr("Replace:   ", strings.Replace("foo", "o", "0", -1))  // Replace:    f00
	ptr("Replace:   ", strings.Replace("foo", "o", "0", 1))   // Replace:    f0o
	ptr("Split:     ", strings.Split("a-b-c-d-e", "-"))       // Split:      [a b c d e]
	ptr("ToLower:   ", strings.ToLower("TEST"))               // ToLower:    test
	ptr("ToUpper:   ", strings.ToUpper("test"))               // ToUpper:    TEST
}
