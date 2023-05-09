package gotutorial

import (
	"os"
	"testing"
)

func TestExit(t *testing.T) {
	// The defers will not be run when using os.Exit
	defer println("defer info")
	os.Exit(0)
}
