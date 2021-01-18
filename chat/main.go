package main

import (
	"flag"
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

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{
		filename: "chat.html",
	}))
	http.Handle("/room", r)

	go r.run()

	log.Println("Web server listening on: ", *address)

	if err := http.ListenAndServe(*address, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
