package main

import (
	"io"
	"net/http"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) { //implementing Handler requires this methoid
	io.WriteString(res, req.RequestURI+"\n") // print request url
	io.WriteString(res, req.URL.Path)
}

func main() {
	var h MyHandler

	http.ListenAndServe(":9000", h)
}
