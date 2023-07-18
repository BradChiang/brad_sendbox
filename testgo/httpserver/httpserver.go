package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v : %v\n", name, h)
		}
	}
}

func contexttest(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: context test handler start")
	defer fmt.Println("server: context test handler end")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "this is context test\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server response : ", err)
		internalerr := http.StatusInternalServerError
		http.Error(w, err.Error(), internalerr)
	}

}

func main() {

	fmt.Println("http server:")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/context", contexttest)

	http.ListenAndServe(":8090", nil)
}
