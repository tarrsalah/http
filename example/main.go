package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world!"))
	})

	log.Println("server listening at http://0.0.0.0:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
