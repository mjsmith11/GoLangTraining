package main

import (
	"fmt"
	"io"
	"net/http"
)

func getAndSetCookie(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("my-cookie")
	fmt.Println(cookie)
	io.WriteString(res, cookie.Value)

	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "COOKIE MONSTER",
	})
}

func main() {
	http.HandleFunc("/", getAndSetCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}
