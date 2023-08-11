package mytesting

import (
	"testing"

	"github.com/arvin-gao/gotutorial/mytesting/mypkg"
)

func TestSetMyMap(t *testing.T) {
	println("1:", mypkg.GetMyMapValue("a"))
	mypkg.SetMyMap("a", 1)
	println("2:", mypkg.GetMyMapValue("a"))
}

func TestGetMyMap(t *testing.T) {
	println("3:", mypkg.GetMyMapValue("a"))
}
