package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", upTown)
	http.HandleFunc("/cat/", youUp)

	http.ListenAndServe(":9000", nil) //nil means use default serve mux (NewServeMux)
}
