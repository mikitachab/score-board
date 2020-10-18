package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}

	s.setupRoutes()
	return s
}

func (s *Server) ListenAndServe(port string) error {
	return http.ListenAndServe(port, s.mux)
}

func (s *Server) setupRoutes() {
	s.mux.Handle("/", s.handleIndex())
}

func compileTemplates() *template.Template {
	templatesPath := "template"
	paths := []string{
		filepath.Join(templatesPath, "base.html"),
		filepath.Join(templatesPath, "head.html"),
		filepath.Join(templatesPath, "navbar.html"),
		filepath.Join(templatesPath, "view.html"),
	}

	templates, err := template.ParseFiles(paths...)
	handleErr(err, "failed to compile templates")

	return templates
}

func (s *Server) handleIndex() http.HandlerFunc {

	// you can put handler setup code here

	// template compiling happens once during setup
	templates := compileTemplates()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			err := templates.ExecuteTemplate(w, "base", nil)
			handleErr(err)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleErr(err error, message ...string) {
	if err != nil {
		if len(message) > 0 {
			err = fmt.Errorf("[%s] -- %w --", message[0], err)
		}
		log.Fatal(err)
	}
}
