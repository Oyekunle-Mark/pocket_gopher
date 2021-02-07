package main

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type poll struct {
	ID      bson.ObjectId  `bson:"_id" json:"id"`
	Title   string         `json:"title"`
	Options []string       `json:"options"`
	Results map[string]int `json:"results,omitempty"`
	APIKey  string         `json:"apiKey"`
}

func (s *Server) handlePolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handlePollsGet(w, r)
		return
	case "POST":
		s.handlePollsPost(w, r)
		return
	case "DELETE":
		s.handlePollsDelete(w, r)
		return
	}

	// not found
	respondHTTPError(w, http.StatusNotFound)
}

func (s *Server) handlePollsGet(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusInternalServerError, errors.New("not implemented"))
}
func (s *Server) handlePollsPost(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusInternalServerError, errors.New("not implemented"))
}
func (s *Server) handlePollsDelete(w http.ResponseWriter,
	r *http.Request) {
	respondError(w, http.StatusInternalServerError, errors.New("not implemented"))
}
