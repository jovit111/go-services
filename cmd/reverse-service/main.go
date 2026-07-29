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
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte(fmt.Sprintf(`{"original":%s,"reversed":%s,"word_count":%d}`, quote(text), quote(string(runes)), strings.Count(string(runes), " ")+1)))
    })
    println("reverse-service listening on :8085")
    http.ListenAndServe(":8085", nil)
}
func quote(s string) string {
    if len(s) > 0 && s[0] != '"' {
        return `"` + s + `"`
    }
    return s
}