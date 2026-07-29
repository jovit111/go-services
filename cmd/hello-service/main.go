package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, `{"message": "Hello, %s!"}`, name)
	})
	fmt.Println("hello-service running on :8080")
	http.ListenAndServe(":8080", nil)
}
