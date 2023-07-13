package gotutorial

import (
	"errors"
	"fmt"
	"testing"
)

var (
	errInvalid = errors.New("invalid")
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg, nil
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func TestError(t *testing.T) {
	ptr(errInvalid)
	for _, i := range []int{7, 42} {
		if r, err := f1(i); err != nil {
			ptr("f1 failed:", err)
		} else {
			ptr("f1 worked:", r)
		}
	}

	_, err := f2(42)
	if myErr, ok := err.(*argError); ok {
		ptr(myErr.arg)
		ptr(myErr.prob)
	}
}

// TODO: more info(errors.Is and errors.As).
// TODO: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
