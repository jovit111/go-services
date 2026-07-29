package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		fmt.Fprintf(w, `{
  "method": "%s",
  "path": "%s",
  "body": %s
}`, r.Method, r.URL.Path, trimQuotes(string(body)))
		w.Header().Set("Content-Type", "application/json")
	})
	fmt.Println("echo-service running on :8081")
	http.ListenAndServe(":8081", nil)
}

func trimQuotes(s string) string {
	if len(s) > 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
