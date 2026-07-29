package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

func main() {
	http.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
		b := make([]byte, 16)
		_, _ = rand.Read(b)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"uuid": "` + hex.EncodeToString(b) + `"}`))
	})
	http.HandleFunc("/api-key", func(w http.ResponseWriter, r *http.Request) {
		b := make([]byte, 32)
		_, _ = rand.Read(b)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"api_key": "` + hex.EncodeToString(b) + `"}`))
	})
	println("generate-service listening on :8084")
	http.ListenAndServe(":8084", nil)
}
