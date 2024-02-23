package packages

import (
	"math/rand"
	"testing"
	"time"
)

// refer to: https://gobyexample.com/random-numbers
func TestRand(t *testing.T) {
	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	ptr()

	ptr(rand.Float64())

	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	ptr()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	ptr()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	ptr()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	p(r3.Intn(100), ",")
	p(r3.Intn(100))
}
