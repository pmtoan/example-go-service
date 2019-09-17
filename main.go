package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Root)
	_ = http.ListenAndServe(":8080", nil)
}

func Root(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("<h1>Hello World!</h1>"))
}
