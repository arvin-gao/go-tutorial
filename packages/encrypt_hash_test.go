package packages

import (
	"crypto/rand"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestSha256(t *testing.T) {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	ptr(s)
	pf("%x\n", bs)
}

func TestBase64(t *testing.T) {
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	ptr(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	ptr(string(sDec))
	ptr()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	ptr(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	ptr(string(uDec))
}

func TestBcrypt(t *testing.T) {
	password := "password"
	password = saltPassword(password)
	hash, err := hashPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Password: %v\nHash: %v\n", password, hash)

	isMatch := checkPasswordHash(password, hash)
	fmt.Println("Match: ", isMatch)
}

func saltPassword(password string) string {
	saltLen := 10
	salt := make([]byte, saltLen)
	rand.Read(salt)
	return password + string(salt)
}

func hashPassword(password string) (string, error) {
	cost := 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
