package main

import (
	"context"
	"flag"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

// Server is the API server.
type Server struct {
	db *mgo.Session
}

func main() {
	address := flag.String("addr", ":8080", "endpoint address")
	mongo := flag.String("mongo", "localhost", "mongodb address")

	log.Println("Dialing mongo", *mongo)

	db, err := mgo.Dial(*mongo)

	if err != nil {
		log.Fatalln("failed to connect to mongo:", err)
	}

	defer db.Close()

	s := &Server{
		db: db,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/polls/", withCORS(withAPIKey(s.handlePolls)))

	log.Println("Starting web server on", *address)

	http.ListenAndServe(*address, mux)
	log.Println("Stopping...")
}

type contextKey struct {
	name string
}

var contextKeyAPIKey = &contextKey{"api-key"}

func APIKey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(contextKeyAPIKey).(string)
	return key, ok
}

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if !isValidAPIKey(key) {
			respondError(w, http.StatusUnauthorized, "invalid API key")
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyAPIKey, key)
		fn(w, r.WithContext(ctx))
	}
}

func isValidAPIKey(key string) bool {
	return key == "abc123"
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}
