package gotutorial

import (
	"fmt"
	"strings"
)

func pass(v ...any) {
}

func ptr(v ...any) {
	fmt.Println(v...)
}

func pTitle(title string) {
	mark := "="
	maxCharLen := 25

	totalLen := len([]rune(title))

	if totalLen > maxCharLen {
		ptr(repeatMark(mark, 2), title, repeatMark(mark, 2))
	}

	markCount := (maxCharLen - totalLen) / 2

	ptr(repeatMark(mark, markCount), title, repeatMark(mark, markCount))
}

func repeatMark(mark string, count int) string {
	return strings.Repeat(mark, count)
}

func pCode(code string) {
	ptr("> " + code)
}

func pfTree(format string, v ...any) {
	format = "\t\t|_ " + format + "\n"
	fmt.Printf(format, v...)
}

func pf(format string, v ...any) {
	fmt.Printf(strings.TrimSpace(format)+"\n", v...)
}

func pSliceLenAndCap(slice []int) {
	pfTree("len(%d), cap(%d)", len(slice), cap(slice))
}

func pPtr(v any) {
	fmt.Printf("%p\n", v)
}
