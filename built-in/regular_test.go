package packages

import (
	"bytes"
	"regexp"
	"testing"
)

func TestRegular(t *testing.T) {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	ptr(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	ptr(r.MatchString("peach"))

	ptr(r.FindString("peach punch"))

	ptr("idx:", r.FindStringIndex("peach punch"))

	ptr(r.FindStringSubmatch("peach punch"))

	ptr(r.FindStringSubmatchIndex("peach punch"))

	ptr(r.FindAllString("peach punch pinch", -1))

	ptr("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	ptr(r.FindAllString("peach punch pinch", 2))

	ptr(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	ptr("regexp:", r)

	ptr(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	ptr(string(out))
}
