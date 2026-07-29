package main
import (
    "fmt"
    "net/http"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name == "" { name = "World" }
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte(fmt.Sprintf(`{"message":"Hello, %s!"}`, name)))
    })
    println("hello-service listening on :8080")
    http.ListenAndServe(":8080", nil)
}