package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		resp := map[string]string{
			"method": r.Method,
			"path":   r.URL.Path,
			"body":   string(body),
		}
		w.Header().Set("Content-Type", "application/json")
		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(resp)
		buf.Truncate(buf.Len() - 1)
		_, _ = w.Write(buf.Bytes())
	})
	println("echo-service listening on :8081")
	http.ListenAndServe(":8081", nil)
}
