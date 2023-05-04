package gotutorial

import (
	"fmt"
	"strings"
)

func pass(v ...any) {
}

func pTitle(title string) {
	mark := "="
	maxCharLen := 50
	tLen := len(title)
	if tLen > maxCharLen {
		println(repeatMark(mark, 2), title, repeatMark(mark, 2))
	}

	markLess := (maxCharLen - tLen) / 2
	println(repeatMark(mark, markLess), title, repeatMark(mark, markLess))
}

func repeatMark(mark string, count int) string {
	return strings.Repeat(mark, count)
}

func pCode(code string) {
	fmt.Println("> " + code)
}

func pfTree(format string, v ...any) {
	format = "\t\t|_ " + format + "\n"
	fmt.Printf(format, v...)
}

func pSliceLenAndCap(slice []int) {
	pfTree("len(%d), cap(%d)", len(slice), cap(slice))
}
