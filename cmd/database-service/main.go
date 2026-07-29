package main
import (
    "database/sql"
    "encoding/json"
    "net/http"
    "sync"
    "time"
    _ "github.com/mattn/go-sqlite3"
)
type Record struct{ ID int `json:"id"`; Name string `json:"name"`; Email string `json:"email"`; CreatedAt time.Time `json:"created_at"` }
var db *sql.DB
var once sync.Once
func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "./app.db")
    if err != nil { panic(err) }
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, email TEXT NOT NULL, created_at datetime)`)
    if err != nil { panic(err) }
}
func main() {
    once.Do(initDB)
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            rows, _ := db.Query("SELECT id, name, email, created_at FROM users")
            var records []Record
            for rows.Next() {
                var rec Record
                _ = rows.Scan(&rec.ID, &rec.Name, &rec.Email, &rec.CreatedAt)
                records = append(records, rec)
            }
            _ = json.NewEncoder(w).Encode(records)
        case http.MethodPost:
            var rec Record
            _ = json.NewDecoder(r.Body).Decode(&rec)
            _, _ = db.Exec("INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)", rec.Name, rec.Email, time.Now())
            _ = json.NewEncoder(w).Encode(map[string]string{"status": "created"})
        default:
            http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
        }
    })
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "db": "sqlite3"})
    })
    println("database-service listening on :9004")
    http.ListenAndServe(":9004", nil)
}