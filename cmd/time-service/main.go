package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		format := r.URL.Query().Get("format")
		now := time.Now().UTC()
		switch format {
		case "unix":
			fmt.Fprintf(w, `{"unix": %d}`, now.Unix())
		case "rfc3339":
			fmt.Fprintf(w, `{"time": "%s"}`, now.Format(time.RFC3339))
		default:
			fmt.Fprintf(w, `{
  "unix": %d,
  "rfc3339": "%s",
  "location": "%s"
}`, now.Unix(), now.Format(time.RFC3339), now.Location())
		}
	})
	fmt.Println("time-service running on :8083")
	http.ListenAndServe(":8083", nil)
}
