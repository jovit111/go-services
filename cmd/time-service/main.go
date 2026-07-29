package main
import (
    "fmt"
    "net/http"
    "time"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        format := r.URL.Query().Get("format")
        now := time.Now().UTC()
        w.Header().Set("Content-Type", "application/json")
        if format == "unix" {
            _, _ = w.Write([]byte(fmt.Sprintf(`{"unix":%d}`, now.Unix())))
        } else if format == "rfc3339" {
            _, _ = w.Write([]byte(fmt.Sprintf(`{"time":"%s"}`, now.Format(time.RFC3339))))
        } else {
            _, _ = w.Write([]byte(fmt.Sprintf(`{"unix":%d,"rfc3339":"%s"}`, now.Unix(), now.Format(time.RFC3339))))
        }
    })
    println("time-service listening on :8083")
    http.ListenAndServe(":8083", nil)
}