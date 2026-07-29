package main
import (
    "encoding/json"
    "net/http"
)
var emails []string
func main() {
    http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
        to := r.URL.Query().Get("to"); subject := r.URL.Query().Get("subject"); body := r.URL.Query().Get("body")
        if to == "" || subject == "" { http.Error(w, `{"error":"to and subject required"}`, http.StatusBadRequest); return }
        emails = append(emails, "To: "+to+"\nSubject: "+subject+"\nBody: "+body)
        _ = json.NewEncoder(w).Encode(map[string]string{"status": "queued", "to": to})
    })
    http.HandleFunc("/inbox", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(emails) })
    http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
        tmpl := map[string]string{"welcome": "Welcome to our service!", "reset": "Click here to reset: /reset", "alert": "New device login detected"}
        _ = json.NewEncoder(w).Encode(tmpl)
    })
    println("email-service listening on :9005")
    http.ListenAndServe(":9005", nil)
}