package main

import (
	"encoding/json"
	"net/http"
	"pocket_gopher/meander"
)

func main() {
	//meander.APIKey = "TODO"
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, meander.Journeys)
	})

	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, data []interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
