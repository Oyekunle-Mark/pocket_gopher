package main

import (
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pocket_gopher/trace"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, r)
}

func main() {
	var address = flag.String("address", ":8080", "The port number of the server")
	flag.Parse()

	gomniauth.SetSecurityKey("take_it.It's all yours.")
	gomniauth.WithProviders(
		github.New(
			"1be45b00bf7deab19191",
			"921c76b75202770aef4006adc1bb85ed756e6027",
			"http://localhost:8080/auth/callback/github",
		),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{
		filename: "chat.html",
	}))
	http.Handle("/login", &templateHandler{
		filename: "login.html",
	})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()

	log.Println("Web server listening on: ", *address)

	if err := http.ListenAndServe(*address, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
