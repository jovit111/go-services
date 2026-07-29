package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("text")
		if text == "" {
			http.Error(w, `{"error":"provide ?text=... param"}`, http.StatusBadRequest)
			return
		}
		runes := []rune(text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		fmt.Fprintf(w, `{
  "original": %s,
  "reversed": %s,
  "word_count": %d
}`, trimQuotes(text), trimQuotes(string(runes)), strings.Count(string(runes), " ")+1)
		w.Header().Set("Content-Type", "application/json")
	})
	fmt.Println("reverse-service running on :8085")
	http.ListenAndServe(":8085", nil)
}

func trimQuotes(s string) string {
	if len(s) > 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}