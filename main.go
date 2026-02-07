package main

import (
	"fmt"
	"net/http"
)

func handle_request(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Sean! (changed by argocd)")
}

func main() {
	fmt.Println("starting server")

	http.HandleFunc("/", handle_request)

	http.ListenAndServe(":8080", nil)
}
