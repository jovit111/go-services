package main
import (
    "encoding/json"
    "net/http"
    "sync"
)
type Config struct{ Key string `json:"key"`; Value string `json:"value"` }
var configs = map[string]string{
    "app.name": "go-services", "app.env": "development", "db.host": "localhost", "db.port": "5432",
    "cache.ttl": "60", "log.level": "info", "api.rate": "100", "queue.name": "default",
    "storage.bucket": "uploads", "email.from": "noreply@example.com",
}
var mu sync.RWMutex
func main() {
    http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")
        mu.RLock(); val, ok := configs[key]; mu.RUnlock()
        if !ok { http.Error(w, `{"error":"not found"}`, http.StatusNotFound); return }
        _ = json.NewEncoder(w).Encode(Config{Key: key, Value: val})
    })
    http.HandleFunc("/configs", func(w http.ResponseWriter, r *http.Request) {
        result := make([]Config, 0, len(configs))
        mu.RLock()
        for k, v := range configs { result = append(result, Config{Key: k, Value: v}) }
        mu.RUnlock()
        _ = json.NewEncoder(w).Encode(result)
    })
    println("config-service listening on :9003")
    http.ListenAndServe(":9003", nil)
}