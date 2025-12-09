package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe("0.0.0.0:3000", http.HandlerFunc(hello))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
