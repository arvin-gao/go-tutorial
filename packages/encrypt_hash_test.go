package packages

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"testing"
)

func TestSha256(t *testing.T) {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	pln(s)
	pf("%x\n", bs)
}

func TestBase64(t *testing.T) {
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	pln(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	pln(string(sDec))
	pln()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	pln(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	pln(string(uDec))
}
