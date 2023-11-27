package implement

import (
	"bufio"
	"fmt"
	"os"
)

func scanInput() {
	in := bufio.NewScanner(os.Stdin)

	// Scan the input line by line
	for in.Scan() {
		in.Text()
	}
	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
