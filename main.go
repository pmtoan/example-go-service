package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Root)
	_ = http.ListenAndServe(":8080", nil)
}

func Root(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		_, _ = rw.Write([]byte("<h1>Hello " + r.URL.Query().Get("name") + "</h1>"))
	} else {
		_, _ = rw.Write([]byte("<h1>Hello World!</h1>"))
	}
}
