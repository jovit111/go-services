package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Job struct {
	ID       string `json:"id"`
	Priority int    `json:"priority"`
	Payload  string `json:"payload"`
}

type Queue struct {
	jobs []Job
	mu   sync.Mutex
}

func (q *Queue) Push(j Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.jobs = append(q.jobs, j)
}

func (q *Queue) Pop() (Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.jobs) == 0 {
		return Job{}, false
	}
	idx := 0
	for i := 1; i < len(q.jobs); i++ {
		if q.jobs[i].Priority > q.jobs[idx].Priority {
			idx = i
		}
	}
	job := q.jobs[idx]
	q.jobs = append(q.jobs[:idx], q.jobs[idx+1:]...)
	return job, true
}

var q = &Queue{}

func main() {
	http.HandleFunc("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		job := Job{ID: r.URL.Query().Get("id"), Payload: r.URL.Query().Get("payload")}
		q.Push(job)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "queued", "id": job.ID})
	})
	http.HandleFunc("/dequeue", func(w http.ResponseWriter, r *http.Request) {
		job, ok := q.Pop()
		if !ok {
			http.Error(w, `{"error":"empty queue"}`, http.StatusNotFound)
			return
		}
		_ = json.NewEncoder(w).Encode(job)
	})
	println("queue-service listening on :9008")
	http.ListenAndServe(":9008", nil)
}
