package main

import "net/http"

func Authorize(w http.ResponseWriter, _ *http.Response) {
	w.Write([]byte("Hello"))
}
