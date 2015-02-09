package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
