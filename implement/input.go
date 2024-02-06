package implement

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Scans the input by fmt package
func scanInputByFmt() {
	var s string

	fmt.Scan(&s)

	fmt.Println("s:", s)
}

// Scans the input by io package.
func scanInputByIO() {
	fmt.Printf("Enter the text:\n")
	writeText, err := os.Open(os.DevNull)
	if err != nil {
		log.Fatalf("failed to open a null device: %s", err)
	}

	defer writeText.Close()
	io.WriteString(writeText, "Write Text")

	readText, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %s", err)
	}

	fmt.Printf("\nLength: %d", len(readText))
	fmt.Printf("\nData Read: \n%s", readText)
}

// Scans the input by bufio package.
func scanInputByBufio() {
	const INPUTDELIMITER = '\n'
	fmt.Print("\nPlease insert a string and press ENTER: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(INPUTDELIMITER)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)

	fmt.Println("Input:", input)
}

// Scans multiple input.
func scanMultiInput() {
	in := bufio.NewScanner(os.Stdin)

	// Scan the input line by line
	for in.Scan() {
		in.Text()
	}
	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
