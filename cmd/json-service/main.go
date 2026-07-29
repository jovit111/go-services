package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func md5Hash(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		raw := r.URL.Query().Get("text")
		if raw == "" {
			http.Error(w, `{"error":"provide ?text=... param"}`, http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"text":   raw,
			"lower":  strings.ToLower(raw),
			"upper":  strings.ToUpper(raw),
			"length": fmt.Sprintf("%d", len(raw)),
			"md5":    md5Hash(raw),
		})
	})
	fmt.Println("json-service running on :8086")
	http.ListenAndServe(":8086", nil)
}
