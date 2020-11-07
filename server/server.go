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
	tl  templateloader.Interface
	db  db.RepositoryInterface
}

// NewServer function create and setup server
func NewServer() *Server {
	return makeServer(templateloader.NewTemplateLoader(), db.GetDB())
}

func makeServer(tl templateloader.Interface, dbrepo db.RepositoryInterface) *Server {
	s := &Server{
		mux: mux.NewRouter(),
		tl:  tl,
		db:  dbrepo,
	}
	s.setupMiddleware()
	s.setupRoutes()
	return s
}

// ListenAndServe starts listening for connection
// and handle them
func (s *Server) ListenAndServe(port string) error {
	return http.ListenAndServe(port, s.mux)
}

func (s *Server) setupMiddleware() {
	s.mux.Use(RecoverMiddleware)
	s.mux.Use(LoggingMiddleware)
}

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("/players", s.handlePlayersList()).Methods("GET")
	s.mux.HandleFunc("/", s.handleIndex()).Methods("GET")
}

func (s *Server) handleIndex() http.HandlerFunc {
	renderIndexTemplate, err := s.tl.GetRenderTemplateFunc("index.html")
	handleErr(err, "failed to setup index template")
	return func(w http.ResponseWriter, r *http.Request) {
		plays := s.db.GetAllPlays()
		var cards []Card

		for _, play := range plays {
			playScores := s.db.GetPlayScoresForPlay(play)
			cards = append(cards, constructCardFromPlayScores(playScores))
		}

		err = renderIndexTemplate(w, cards)
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
