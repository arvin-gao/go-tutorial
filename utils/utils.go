package utils

import "fmt"

// CheckError use `CheckErrorWithMessage(err, "")`.
func CheckError(err error) {
	CheckErrorWithMessage(err, "")
}

// CheckErrorWithMessage if err is not nil, then panic the error(err) with the message(msg).
func CheckErrorWithMessage(err error, msg string) {
	if err != nil {
		if msg == "" {
			panic(err)
		} else {
			panic(fmt.Sprintf("%s, Msg:%s", err, msg))
		}
	}
}
