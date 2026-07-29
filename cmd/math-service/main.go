package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		a, _ := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
		b, _ := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
		fmt.Fprintf(w, `{"result": %g}`, a+b)
		w.Header().Set("Content-Type", "application/json")
	})
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		a, _ := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
		b, _ := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
		fmt.Fprintf(w, `{"result": %g}`, a*b)
		w.Header().Set("Content-Type", "application/json")
	})
	http.HandleFunc("/factorial", func(w http.ResponseWriter, r *http.Request) {
		n, _ := strconv.Atoi(r.URL.Query().Get("n"))
		fmt.Fprintf(w, `{"result": %d}`, factorial(n))
		w.Header().Set("Content-Type", "application/json")
	})
	fmt.Println("math-service running on :8087")
	http.ListenAndServe(":8087", nil)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
