package main

import (
	"encoding/json"
	"net/http"
)

var files = map[string][]byte{}

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		content := []byte(r.URL.Query().Get("content"))
		if name == "" {
			http.Error(w, `{"error":"name required"}`, http.StatusBadRequest)
			return
		}
		files[name] = content
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "saved", "name": name})
	})
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		val, ok := files[r.URL.Query().Get("name")]
		if !ok {
			http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(val)
	})
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		names := make([]string, 0, len(files))
		for k := range files {
			names = append(names, k)
		}
		_ = json.NewEncoder(w).Encode(names)
	})
	println("file-service listening on :9006")
	http.ListenAndServe(":9006", nil)
}
