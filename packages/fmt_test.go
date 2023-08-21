package packages

import (
	"fmt"
	"os"
	"testing"
)

// refer to: https://gobyexample.com/string-formatting
func TestFmt(t *testing.T) {

	type point struct {
		x, y int
	}

	p := point{1, 2}

	ptrf("struct1: %v\n", p)                     //  {1 2}
	ptrf("struct2: %+v\n", p)                    //  {x:1 y:2}
	ptrf("struct3: %#v\n", p)                    //  packages.point{x:1, y:2}
	ptrf("type: %T\n", p)                        //  packages.point
	ptrf("bool: %t\n", true)                     //  true
	ptrf("int: %d\n", 123)                       //  123
	ptrf("bin: %b\n", 14)                        //  1110
	ptrf("char: %c\n", 33)                       //  !
	ptrf("hex: %x\n", 456)                       //  1c8
	ptrf("float1: %f\n", 78.9)                   //  78.900000
	ptrf("float2: %e\n", 123400000.0)            //  1.234000e+08
	ptrf("float3: %E\n", 123400000.0)            //  1.234000E+08
	ptrf("str1: %s\n", "\"string\"")             //  "string"
	ptrf("str2: %q\n", "\"string\"")             //  "\"string\""
	ptrf("str3: %x\n", "hex this")               //  6865782074686973
	ptrf("pointer: %p\n", &p)                    //  0x140000ac190
	ptrf("width1: |%6d|%6d|\n", 12, 345)         //  |    12|   345|
	ptrf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)   //  |  1.20|  3.45|
	ptrf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) //  |1.20  |3.45  |
	ptrf("width4: |%6s|%6s|\n", "foo", "b")      //  |   foo|     b|
	ptrf("width5: |%-6s|%-6s|\n", "foo", "b")    //  |foo   |b     |

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s) //  a string

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") //  an error
}
