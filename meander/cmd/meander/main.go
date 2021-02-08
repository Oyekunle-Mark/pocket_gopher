package main

import (
	"encoding/json"
	"net/http"
	"pocket_gopher/meander"
	"strconv"
	"strings"
)

func main() {
	meander.APIKey = "Some cryptic key"

	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, meander.Journeys)
	})
	http.HandleFunc("/recommendations", cors(func(w http.ResponseWriter, r *http.Request) {
		q := &meander.Query{
			Journey: strings.Split(r.URL.Query().Get("journey"), "|"),
		}

		var err error

		q.Lat, err = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.Lng, err = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.Radius, err = strconv.Atoi(r.URL.Query().Get("radius"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		q.CostRangeStr = r.URL.Query().Get("cost")

		places := q.Run()
		respond(w, places)
	}))

	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func cors(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}

func respond(w http.ResponseWriter, data []interface{}) error {
	publicData := make([]interface{}, len(data))

	for i, d := range data {
		publicData[i] = meander.Public(d)
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(publicData)
}
