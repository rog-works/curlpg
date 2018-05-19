package main

import (
	"net/http"

	"./server"
)

func main() {
	http.HandleFunc("/", server.Handler)
	http.ListenAndServe(":80", nil)
}
