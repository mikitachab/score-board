package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mikitachab/score-board/db"
	"github.com/mikitachab/score-board/templateloader"
)

// Server is a main application abstraction
type Server struct {
	mux *mux.Router
	tl  *templateloader.TemplateLoader
	db  *db.Repository
}

// NewServer function create and setup server
func NewServer() *Server {
	s := &Server{
		mux: mux.NewRouter(),
		tl:  templateloader.NewTemplateLoader(),
		db:  db.GetDB(),
	}

	s.setupRoutes()
	return s
}

// ListenAndServe starts listening for connection
// and handle them
func (s *Server) ListenAndServe(port string) error {
	return http.ListenAndServe(port, s.mux)
}

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("/players", s.handlePlayersList()).Methods("GET")
	s.mux.HandleFunc("/", s.handleIndex()).Methods("GET")
}

func (s *Server) handleIndex() http.HandlerFunc {
	renderIndexTemplate, err := s.tl.GetRenderTemplateFunc("index.html")
	handleErr(err, "failed to setup index template")

	return func(w http.ResponseWriter, r *http.Request) {
		err := renderIndexTemplate(w, nil)
		handleErr(err)
	}
}

func (s *Server) handlePlayersList() http.HandlerFunc {
	renderPlayersListTemplate, err := s.tl.GetRenderTemplateFunc("players_list.html")
	handleErr(err, "failed to setup player_list template")

	return func(w http.ResponseWriter, r *http.Request) {
		players := s.db.GetAllPlayers()
		err := renderPlayersListTemplate(w, players)
		handleErr(err)
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
