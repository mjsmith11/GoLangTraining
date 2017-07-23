package main

import (
	"io"
	"net/http"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

func youUp(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty catty catty")
}

func main() {
	mux := http.NewServeMux()
	//will match the most specific
	// /hello matches upTown
	// /cat/hello matches youUp
	mux.HandleFunc("/", upTown)
	mux.HandleFunc("/cat/", youUp)

	http.ListenAndServe(":9000", mux)

}
