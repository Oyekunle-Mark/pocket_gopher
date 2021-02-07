package main

import (
	"errors"
	"gopkg.in/mgo.v2"
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
	session := s.db.Copy()
	defer session.Close()

	c := session.DB("ballots").C("polls")
	var q *mgo.Query
	p := NewPath(r.URL.Path)

	if p.HasID() {
		q = c.FindId(bson.ObjectIdHex(p.ID))
	} else {
		q = c.Find(nil)
	}

	var result []*poll

	if err := q.All(&result); err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respond(w, http.StatusOK, &result)
}

func (s *Server) handlePollsPost(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusInternalServerError, errors.New("not implemented"))
}

func (s *Server) handlePollsDelete(w http.ResponseWriter,
	r *http.Request) {
	respondError(w, http.StatusInternalServerError, errors.New("not implemented"))
}