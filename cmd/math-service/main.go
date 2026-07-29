package main
import (
    "fmt"
    "net/http"
    "strconv"
)
func factorial(n int) int {
    if n <= 1 { return 1 }
    return n * factorial(n-1)
}
func main() {
    http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
        a, _ := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
        b, _ := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte(fmt.Sprintf(`{"result":"%g"}`, a+b)))
    })
    http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
        a, _ := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
        b, _ := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte(fmt.Sprintf(`{"result":"%g"}`, a*b)))
    })
    http.HandleFunc("/factorial", func(w http.ResponseWriter, r *http.Request) {
        n, _ := strconv.Atoi(r.URL.Query().Get("n"))
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte(fmt.Sprintf(`{"result":"%d"}`, factorial(n))))
    })
    fmt.Println("math-service running on :8087")
    http.ListenAndServe(":8087", nil)
}