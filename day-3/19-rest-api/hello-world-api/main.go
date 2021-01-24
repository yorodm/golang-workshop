package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
