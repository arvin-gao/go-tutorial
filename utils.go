package gotutorial

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var ptrPrefix = ""

func pass(v ...any) {
	// nothing to do here.
}

func ptr(v ...any) {
	fmt.Print(ptrPrefix)
	fmt.Println(v...)
}

// ptrf() is an alias for fmt.Printf() function.
func ptrf(format string, v ...any) {
	fmt.Printf(ptrPrefix+strings.TrimSpace(format)+"\n", v...)
}

func pPtr(v any) {
	fmt.Printf("%p\n", v)
}

func pTitle(title string) {
	ptrPrefix = ""
	mark := "="
	maxCharLen := 25

	totalLen := utf8.RuneCountInString(title)

	if totalLen > maxCharLen {
		ptr(repeatMark(mark, 2), title, repeatMark(mark, 2))
		return
	}

	markCount := (maxCharLen - totalLen) / 2

	ptr(repeatMark(mark, markCount), title, repeatMark(mark, markCount))
}

func repeatMark(mark string, count int) string {
	return strings.Repeat(mark, count)
}

func ptrSubject(title any) {
	ptrPrefix = "* "
	ptr(title)
	ptrPrefix = strings.Repeat(" ", 2)
}

func pCode(code string) {
	ptr("> " + code)
}

func pCode2(code string) {
	str := strings.Split(code, "\n")
	ptr("[Code]:")
	space := strings.Repeat(" ", 2)
	for _, s := range str {
		print(space)
		ptr(s)
	}
}

func ptrfTree(format string, v ...any) {
	format = "\t\t|_ " + format + "\n"
	fmt.Printf(format, v...)
}

func ptrSliceLenAndCap(slice []int) {
	ptrf("len: %d", len(slice))
	ptrf("cap: %d", cap(slice))
}
