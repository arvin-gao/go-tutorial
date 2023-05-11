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

func TestStringType(t *testing.T) {
	// 修改字串 by index
	str := "abc"
	c := []byte(str)
	println("origin string:", str)

	c[0] = 'c'

	pCode("c[0] = 'c'")
	pfTree("%s", c)

	// 使用 for range
	pTitle("for range the string")
	pCode("for i, v := range \"abc\"")
	for i, v := range "abc" {
		println(i, v)
	}

	// 計算長度
	pTitle("string length")
	str = "你好 abc"
	pCode("str = \"你好 abc\"")
	pfTree("len(str): %d", len(str))
	pfTree("len([]rune(str)): %d", len([]rune(str)))
	pfTree("utf8.RuneCountInString(str): %d", utf8.RuneCountInString(str))

	// 連接字串
	pTitle("string concatenation")
	pCode("str =\"\"")
	pCode("str +=\"abc\"")
	str = ""
	str += "abc"
	pfTree("str: %s", str)
	var strBuf strings.Builder
	pCode("var strBuf strings.Builder")
	pCode("strBuf.WriteString(\"test\")")
	strBuf.WriteString("test")
	pfTree("strBuf.String(): %s", strBuf.String())

	// 地址變化
	addWord := func(s *string, count int) {
		*s += strings.Repeat("a", count)
	}
	addWordToBuf := func(buf *strings.Builder, count int) {
		buf.WriteString(strings.Repeat("a", count))
	}
	printStringLenAndCap := func(str *string) {
		pfTree("len: %d, cap: %d", len(*str), cap([]byte(*str)))
	}
	printBufLenAndCap := func(buf *strings.Builder) {
		pfTree("len: %d, cap: %d", strBuf.Len(), strBuf.Cap())
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
		pf("%#U starts at %d, size: %d\n", runeV, i, size)
		w = size
	}
}
