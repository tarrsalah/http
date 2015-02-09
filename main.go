package main

import (
	"fmt"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello  Watch!")
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":80", nil)
}
