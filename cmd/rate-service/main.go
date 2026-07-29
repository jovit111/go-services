package main
import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
    "time"
)
type Limit struct{ Key string `json:"key"`; Remaining int `json:"remaining"`; Reset int64 `json:"reset_at"` }
type Limiter struct{ limits map[string]*Limit; mu sync.Mutex }
func (l *Limiter) Check(key string, limit int) bool {
    l.mu.Lock(); defer l.mu.Unlock()
    now := time.Now().Unix()
    if _, exists := l.limits[key]; !exists || now > l.limits[key].Reset {
        l.limits[key] = &Limit{Key: key, Remaining: limit - 1, Reset: now + 60}
        return true
    }
    l.limits[key].Remaining--
    return l.limits[key].Remaining >= 0
}
var limiter = &Limiter{limits: make(map[string]*Limit)}
func main() {
    http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")
        limit := 10
        fmt.Sscanf(r.URL.Query().Get("limit"), "%d", &limit)
        allowed := limiter.Check(key, limit)
        _ = json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "allowed": allowed, "reset_at": time.Now().Add(time.Minute).Unix()})
    })
    println("rate-service listening on :9009")
    http.ListenAndServe(":9009", nil)
}