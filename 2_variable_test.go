package gotutorial

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

type (
	Point struct{ x, y float64 } // Point and struct{ x, y float64 } are different types
	polar Point                  // polar and Point denote different types
)

func TestDeclaration(t *testing.T) {
	var a int
	var b int = 1
	var c = 1
	var d *int
	e := 1

	var a1, b1, c1 = 1, 2.0, "3"

	var (
		a2 int
		b2 int
	)

	v1, v2, err := getMultiVars()
	_, v22, _ := getMultiVars()

	const cInt = 1
	const cStr string = "1"

	pass(a2, b2, a, b, c, d, cInt, cStr, v1, v2, v22, err, a1, b1, c1, e)
}

func TestInitVar(t *testing.T) {
	var ch chan string
	fmt.Println(ch == nil)
}

func TestConvertType(t *testing.T) {
	// numbers
	var num = 1
	numA := int64(num)
	numB := float64(num)
	_ = numA + int64(numB)

	// string.
	str := string([]byte("abc"))
	bys := []byte("abc")
	char := str[0] // byte
	str = strconv.Itoa(123)
	_ = str + str + string(char)

	// array & slice
	var arr [1]int
	slice := arr[:]
	arr2 := [1]int(slice)
	_, _ = arr[0], slice[0]

	// interface | any
	var i any
	i = bys
	_bys, _ := i.([]byte)
	i = arr2
	_arr2 := i.([1]int)

	pass(_bys, _arr2)
}

func getMultiVars() (int, string, error) {
	return 0, "", errors.New("err")
}

func TestStringType(t *testing.T) {
	// 修改字串 by index
	pTitle("modify string")
	str := "abc"
	c := []byte(str)
	println("origin string:", str)
	pCode("c[0] = 'c'")
	c[0] = 'c'
	pfTree("%s", c)

	// 使用 for range
	pTitle("for range the string")
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
	// TODO: string cap?(memory, copy).
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
