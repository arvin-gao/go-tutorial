package gotutorial

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestCommonPkg(t *testing.T) {
	var (
		str string
		// num int
		// err error
	)

	println("")
	println(fmt.Sprintf("%d, %d", 1, 2))

	// strings.
	strSlice := strings.Fields(" dd dte aseg ")
	str = strings.Join(strSlice, ",")
	fmt.Printf("%#v\n", strSlice)
	fmt.Printf("%s\n", str)

	// strconv
	str = strconv.Itoa(10)
	// num, _ = strconv.Atoi(str)

}
