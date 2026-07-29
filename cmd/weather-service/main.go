package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Location string `json:"location"`
	TempC    int    `json:"temp_c"`
	TempF    int    `json:"temp_f"`
	Summary  string `json:"summary"`
	Updated  string `json:"updated_at"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		location := r.URL.Query().Get("location")
		if location == "" {
			location = "Unknown"
		}
		_ = json.NewEncoder(w).Encode(WeatherResponse{
			Location: location,
			TempC:    22,
			TempF:    72,
			Summary:  "Partly cloudy",
			Updated:  time.Now().UTC().Format(time.RFC3339),
		})
	})
	fmt.Println("weather-service running on :8088")
	http.ListenAndServe(":8088", nil)
}
