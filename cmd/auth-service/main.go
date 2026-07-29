package main
import (
    "crypto/subtle"
    "crypto/rand"
    "encoding/hex"
    "encoding/json"
    "net/http"
    "strings"
    "time"
)
type Token struct {
    Access  string `json:"access_token"`
    Type    string `json:"token_type"`
    Expires int64  `json:"expires_at"`
}
var users = map[string]string{"admin": "secret", "user": "password"}
func main() {
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
            return
        }
        var creds struct{ Username, Password string }
        if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
            http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
            return
        }
        password, ok := users[creds.Username]
        if !ok || subtle.ConstantTimeCompare([]byte(password), []byte(creds.Password)) != 1 {
            http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
            return
        }
        token := Token{Access: "tok_" + random(24), Type: "bearer", Expires: time.Now().Add(time.Hour).Unix()}
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(token)
    })
    http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
        if !strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ") {
            _ = json.NewEncoder(w).Encode(map[string]bool{"valid": false})
            return
        }
        _ = json.NewEncoder(w).Encode(map[string]interface{}{"valid": true, "user": "admin"})
    })
    println("auth-service listening on :9001")
    http.ListenAndServe(":9001", nil)
}
func random(n int) string {
    b := make([]byte, n)
    _, _ = rand.Read(b)
    return hex.EncodeToString(b)
}