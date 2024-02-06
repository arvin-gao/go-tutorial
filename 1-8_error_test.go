package gotutorial

import (
	"errors"
	"fmt"
	"testing"
)

var (
	errInvalid = errors.New("invalid message")
)

type myErrorStruct struct {
	msg string
}

func (e *myErrorStruct) Error() string {
	return fmt.Sprintf("This is from my error struct. Param message:%s", e.msg)
}

func TestError(t *testing.T) {
	f := func(n int) error {
		switch n {
		case 1:
			return errors.New("1 error")
		case 2:
			return errInvalid
		case 3:
			return &myErrorStruct{"can't work with it"}
		default:
			return nil
		}
	}

	nums := []int{1, 2, 3, 4}
	for _, n := range nums {
		if err := f(n); err != nil {
			ptr(err)
		}
	}

	err := f(3)
	if myErr, ok := err.(*myErrorStruct); ok {
		ptr(myErr.msg)
	}
}
