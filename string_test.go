package gotutorial

import (
	"testing"
	"unicode/utf8"
)

// string: Go string literals are UTF-8 encoded text.

// Since strings are equivalent to []byte,
// this will produce the length of the raw bytes stored within

func TestString(t *testing.T) {
	s := "Hello, 世界"
	println("string length:", len(s))

	b := []byte(s)
	println("[]byte length:", len(b))

	r := []rune(s)
	println("[]rune length:", len(r))

	println("utf8.RuneCountInString(string):", utf8.RuneCountInString(s))

	print("loop with index: ")
	for i := 0; i < len(s); i++ {
		print(s[i], " ")
	}

	println()
	print("loop with range: ")
	for _, v := range s {
		print(v, " ")
	}
}

func TestStringWithDecodeRune(t *testing.T) {
	s := "hi 你好"
	for i, w := 0, 0; i < len(s); i += w {
		runeV, size := utf8.DecodeRuneInString(s[i:])
		pf("%#U starts at %d, size: %d\n", runeV, i, size)
		w = size
	}
}
