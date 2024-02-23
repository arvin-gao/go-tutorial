package packages

import (
	"encoding/base64"
	"testing"
)

func TestBase64URL(t *testing.T) {
	src := []byte("abc123!?$*&()'-=@~")

	println(base64.URLEncoding.EncodeToString(src))
	println(base64.StdEncoding.EncodeToString(src))
}
