package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/forward", func(w http.ResponseWriter, r *http.Request) {
		target := r.URL.Query().Get("target")
		if target == "" {
			http.Error(w, `{"error":"target required"}`, http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]string{
			"method": r.Method, "target": target, "status": "forwarded",
		})
	})
	http.HandleFunc("/routes", func(w http.ResponseWriter, r *http.Request) {
		routes := []map[string]string{
			{"path": "/api/v1", "target": "http://api:8080"},
			{"path": "/api/v2", "target": "http://api2:8081"},
			{"path": "/health", "target": "http://health:8082"},
			{"path": "/static", "target": "http://static:9006"},
		}
		_ = json.NewEncoder(w).Encode(routes)
	})
	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		sanitized := make(map[string]string)
		for k, v := range r.Header {
			sanitized[k] = strings.Join(v, ", ")
		}
		_ = json.NewEncoder(w).Encode(sanitized)
	})
	println("proxy-service listening on :9007")
	http.ListenAndServe(":9007", nil)
}
