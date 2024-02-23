package packages

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
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

func TestContextWithValue(t *testing.T) {
	type ctxKey int

	var (
		traceIdKey ctxKey = 0
		userIdKey  ctxKey = 1
	)

	tranceId := uuid.New().String()
	ctx := context.WithValue(context.Background(), traceIdKey, tranceId)

	// External service.
	getUserIdService := func(ctx context.Context) context.Context {
		id := rand.Int63n(1000)
		log.Printf("get user id %v. trace ID: %v", id, ctx.Value(traceIdKey))
		return context.WithValue(ctx, userIdKey, id)
	}

	ctx = getUserIdService(ctx)

	log.Printf("finish to doing the test function service. trace ID: %v, user ID: %v", ctx.Value(traceIdKey), ctx.Value(userIdKey))
}

func TestContextWithCancelOnMultiGoroutine(t *testing.T) {
	doSystemA := func(ctx context.Context) error {
		time.Sleep(100 * time.Millisecond)
		return errors.New("system-A failed")
	}

	doSystemB := func(ctx context.Context) {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("system-B process finished")
		case <-ctx.Done():
			fmt.Println("system-B process cancelled")
		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		if err := doSystemA(ctx); err != nil {
			cancel()
		}
	}()

	doSystemB(ctx)
}

func TestContextWithTimeout(t *testing.T) {
	timeout := 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	requestResource := func(ctx context.Context) (any, error) {
		select {
		case <-time.After(200 * time.Millisecond):
			fmt.Println("request success")
			return nil, nil
		case <-ctx.Done():
			return nil, errors.New("request process cancelled")
		}
	}

	source, err := requestResource(ctx)
	if err != nil {
		log.Println("error:", err)
		return
	}
	fmt.Println(source)
}
