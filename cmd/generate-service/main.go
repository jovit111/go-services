package main
import (
    "crypto/rand"
    "encoding/hex"
    "net/http"
)
func random(n int) string {
    b := make([]byte, n)
    _, _ = rand.Read(b)
    return hex.EncodeToString(b)
}
func main() {
    http.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte(`{"uuid":"` + random(16) + `"}`))
    })
    http.HandleFunc("/api-key", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte(`{"api_key":"` + random(32) + `"}`))
    })
    println("generate-service listening on :8084")
    http.ListenAndServe(":8084", nil)
}