package main
import (
    "encoding/json"
    "net/http"
    "strconv"
    "sync"
    "time"
)
type Item struct{ Key, Value string; Expires int64 `json:"expires_at"` }
type Cache struct{ data map[string]Item; mu sync.RWMutex }
func (c *Cache) Set(key, value string, ttl int64) {
    c.mu.Lock(); defer c.mu.Unlock()
    c.data[key] = Item{Key: key, Value: value, Expires: time.Now().Unix() + ttl}
}
func (c *Cache) Get(key string) (Item, bool) {
    c.mu.RLock(); defer c.mu.RUnlock()
    item, ok := c.data[key]
    if !ok || time.Now().Unix() > item.Expires { return Item{}, false }
    return item, true
}
var store = &Cache{data: make(map[string]Item)}
func main() {
    http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")
        val := r.URL.Query().Get("value")
        ttl, _ := strconv.ParseInt(r.URL.Query().Get("ttl"), 10, 64)
        if ttl == 0 { ttl = 60 }
        store.Set(key, val, ttl)
        _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })
    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
        item, ok := store.Get(r.URL.Query().Get("key"))
        if !ok { http.Error(w, `{"error":"not found"}`, http.StatusNotFound); return }
        _ = json.NewEncoder(w).Encode(item)
    })
    println("cache-service listening on :9002")
    http.ListenAndServe(":9002", nil)
}