package mytesting

import (
	"testing"

	"github.com/arvin-gao/gotutorial/testing/temp"
)

func TestSetMyMap(t *testing.T) {
	println("1:", temp.GetMyMapValue("a"))
	temp.SetMyMap("a", 1)
	println("2:", temp.GetMyMapValue("a"))
}

func TestGetMyMap(t *testing.T) {
	println("3:", temp.GetMyMapValue("a"))
}
