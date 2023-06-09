package packages

import (
	"bufio"
	"fmt"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	pln("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		pln(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func TestHttpServer(t *testing.T) {
	hello := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
	}

	headers := func(w http.ResponseWriter, req *http.Request) {

		for name, headers := range req.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
