package main

import (
	"encoding/json"
	"net/http"
	"pocket_gopher/meander"
)

func main() {
	meander.APIKey = "Some cryptic key"

	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, meander.Journeys)
	})

	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, data []interface{}) error {
	publicData := make([]interface{}, len(data))

	for i, d := range data {
		publicData[i] = meander.Public(d)
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(publicData)
}
