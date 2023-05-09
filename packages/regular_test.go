package packages

import (
	"bytes"
	"regexp"
	"testing"
)

func TestRegular(t *testing.T) {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	pln(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	pln(r.MatchString("peach"))

	pln(r.FindString("peach punch"))

	pln("idx:", r.FindStringIndex("peach punch"))

	pln(r.FindStringSubmatch("peach punch"))

	pln(r.FindStringSubmatchIndex("peach punch"))

	pln(r.FindAllString("peach punch pinch", -1))

	pln("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	pln(r.FindAllString("peach punch pinch", 2))

	pln(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	pln("regexp:", r)

	pln(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	pln(string(out))
}
