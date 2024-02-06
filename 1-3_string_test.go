package gotutorial

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// string: Go string literals are UTF-8 encoded text.

// Since strings are equivalent to []byte,
// this will produce the length of the raw bytes stored within

func TestString(t *testing.T) {
	s := "Hello, 世界"
	ptr("string length:", len(s))

	b := []byte(s)
	ptr("[]byte length:", len(b))

	r := []rune(s)
	ptr("[]rune length:", len(r))

	ptr("utf8.RuneCountInString(string):", utf8.RuneCountInString(s))

	print("loop with index: ")
	for i := 0; i < len(s); i++ {
		print(s[i], " ") // byte
	}

	ptr()
	print("loop with range: ")
	for _, v := range s {
		print(v, " ") // rune
	}
}

func TestStringType(t *testing.T) {
	// Modify the string by index
	str := "abc"
	c := []byte(str) // []byte{'c', 'b', 'c'}

	c[0] = 'c'
	ptr(string(c))

	// Ranges the string.
	for i, v := range "abc" {
		ptr(i, v)
	}

	// String length.
	str = "中文 abc"
	ptrlnf("len(str): %d", len(str))
	ptrlnf("len([]rune(str)): %d", len([]rune(str)))
	ptrlnf("utf8.RuneCountInString(str): %d", utf8.RuneCountInString(str))

	// String concatenation.
	pTitle("String concatenation")
	str = ""
	str += "abc"
	ptrfTree("str: %s", str)
	var strBuf strings.Builder
	pCode("var strBuf strings.Builder")
	pCode("strBuf.WriteString(\"test\")")
	strBuf.WriteString("test")
	ptrfTree("strBuf.String(): %s", strBuf.String())

	// TODO: wait
	// 地址變化
	addWord := func(s *string, count int) {
		*s += strings.Repeat("a", count)
	}
	addWordToBuf := func(buf *strings.Builder, count int) {
		buf.WriteString(strings.Repeat("a", count))
	}
	printStringLenAndCap := func(str *string) {
		ptrfTree("len: %d, cap: %d", len(*str), cap([]byte(*str)))
	}
	printBufLenAndCap := func(buf *strings.Builder) {
		ptrfTree("len: %d, cap: %d", strBuf.Len(), strBuf.Cap())
	}
	pTitle("string growing")

	str = ""
	addWord(&str, 1)
	printStringLenAndCap(&str)
	addWord(&str, 31)
	printStringLenAndCap(&str)
	addWord(&str, 1)
	printStringLenAndCap(&str)

	pTitle("string buffer growing")
	strBuf = strings.Builder{}
	pCode("strBuf = strings.Builder{}")
	printBufLenAndCap(&strBuf)
	pCode("add 1 string")
	addWordToBuf(&strBuf, 1)
	printBufLenAndCap(&strBuf)
	pCode("add 9 string")
	addWordToBuf(&strBuf, 9)
	printBufLenAndCap(&strBuf)
	pCode("add 7 string")
	addWordToBuf(&strBuf, 7)
	printBufLenAndCap(&strBuf)
	pCode("add 16 string")
	addWordToBuf(&strBuf, 16)
	printBufLenAndCap(&strBuf)
	pCode("strBuf.Grow(1000)")
	strBuf.Grow(1000)
	printBufLenAndCap(&strBuf)
	pCode("add 1000 string")
	addWordToBuf(&strBuf, 1000)
	printBufLenAndCap(&strBuf)
}

func TestStringWithDecodeRune(t *testing.T) {
	s := "hi 你好"
	for i, w := 0, 0; i < len(s); i += w {
		runeV, size := utf8.DecodeRuneInString(s[i:])
		ptrlnf("%#U starts at %d, size: %d\n", runeV, i, size)
		w = size
	}
}
