package main
import (
    "encoding/json"
    "net/http"
    "time"
)
type Health struct {
    Status    string            `json:"status"`
    Timestamp string            `json:"timestamp"`
    Version   string            `json:"version"`
    Checks    map[string]string `json:"checks"`
}
func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(Health{Status:"healthy", Timestamp:time.Now().UTC().Format(time.RFC3339), Version:"1.0.0", Checks:map[string]string{"database":"ok","cache":"ok","queue":"ok"}})
    })
    http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ready"))
    })
    println("health-service listening on :8082")
    http.ListenAndServe(":8082", nil)
}