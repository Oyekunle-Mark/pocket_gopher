package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type authHandler struct {
	next http.Handler
}

func (a *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")

	if err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	a.next.ServeHTTP(w, r)
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{
		next: handler,
	}
}

// loginHandler handles the third-party login process.
// format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")

	action := segments[2]
	provider := segments[3]

	switch action {
	case "login":
		log.Println("TODO handle login for", provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
