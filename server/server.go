package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mikitachab/score-board/templateloader"
)

type Server struct {
	mux *http.ServeMux
	tl  *templateloader.TemplateLoader
}

func NewServer() *Server {
	s := &Server{
		mux: http.NewServeMux(),
		tl:  templateloader.NewTemplateLoader(),
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

func (s *Server) handleIndex() http.HandlerFunc {
	renderIndexTemplate, err := s.tl.GetRenderTemplateFunc("index.html")
	handleErr(err, "failed to setup index template")

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			err := renderIndexTemplate(w, nil)
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
