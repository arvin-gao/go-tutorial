package gotutorial

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// TODO: https://gobyexample.com/context.
func TestContext(t *testing.T) {
	hello := func(w http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		println("server: hello handler started")
		defer println("server: hello handler ended")

		select {
		case <-time.After(10 * time.Second):
			fmt.Fprintf(w, "hello\n")
		case <-ctx.Done():

			err := ctx.Err()
			println("server:", err)
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)
		}
	}

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
