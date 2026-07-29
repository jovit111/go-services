package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type LogEntry struct {
	Timestamp string            `json:"timestamp"`
	Method    string            `json:"method"`
	Path      string            `json:"path"`
	Headers   map[string]string `json:"headers"`
	IP        string            `json:"ip"`
}

func main() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		headers := make(map[string]string)
		for k, v := range r.Header {
			headers[k] = strings.Join(v, ", ")
		}
		entry := LogEntry{Timestamp: time.Now().UTC().Format(time.RFC3339Nano), Method: r.Method, Path: r.URL.Path, Headers: headers, IP: r.RemoteAddr}
		fmt.Printf("%+v\n", entry)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(entry)
	})
	println("logger-service listening on :8089")
	http.ListenAndServe(":8089", nil)
}
