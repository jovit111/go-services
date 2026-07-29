package main
import (
    "encoding/json"
    "fmt"
    "net/http"
)
var objects = map[string][]byte{}
func main() {
    http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")
        val := []byte(r.URL.Query().Get("value"))
        if key == "" { http.Error(w, `{"error":"key required"}`, http.StatusBadRequest); return }
        objects[key] = val
        _ = json.NewEncoder(w).Encode(map[string]string{"status": "stored", "key": key, "size": fmt.Sprintf("%d", len(val))})
    })
    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
        val, ok := objects[r.URL.Query().Get("key")]
        if !ok { http.Error(w, `{"error":"not found"}`, http.StatusNotFound); return }
        w.Header().Set("Content-Type", "application/octet-stream")
        _, _ = w.Write(val)
    })
    http.HandleFunc("/keys", func(w http.ResponseWriter, r *http.Request) {
        keys := make([]string, 0, len(objects))
        for k := range objects { keys = append(keys, k) }
        _ = json.NewEncoder(w).Encode(keys)
    })
    println("storage-service listening on :9010")
    http.ListenAndServe(":9010", nil)
}