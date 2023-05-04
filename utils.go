package gotutorial

import (
	"fmt"
	"strings"
)

func pass(v ...any) {
}

func pTitle(title string) {
	mark := "="
	maxCharLen := 40

	var totalLen int
	for _, v := range title {
		totalLen += 1
		if v > 'z' {
			totalLen += 1
		}
	}

	if totalLen > maxCharLen {
		println(repeatMark(mark, 2), title, repeatMark(mark, 2))
	}

	markCount := (maxCharLen - totalLen) / 2

	println(repeatMark(mark, markCount), title, repeatMark(mark, markCount))
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

func pPtr(v any) {
	fmt.Printf("%p\n", v)
}
