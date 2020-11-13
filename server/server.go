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

// ListenAndServe starts listening for connection
// and handle them
func (s *Server) ListenAndServe(port string) error {
	return http.ListenAndServe(port, s.mux)
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

func (s *Server) setupMiddleware() {
	s.mux.Use(RecoverMiddleware)
	s.mux.Use(LoggingMiddleware)
}

// HandlerCtx server context
// to be passed to TemplateHandlerFunc
type HandlerCtx struct {
	DB       db.RepositoryInterface
	Template templateloader.TemplateInterface
}

func (s *Server) makeHandlerCtx(templateName string) *HandlerCtx {
	template, err := s.tl.LoadTemplate(templateName)
	handleErr(err, fmt.Sprintf("failed to setup %s template", templateName))
	return &HandlerCtx{s.db, template}
}

// View is a func that return http handler
// that should render template
type View func(*HandlerCtx) http.HandlerFunc

// Handle is function to connect url pattern, template and handler function
func (s *Server) Handle(pattern, templateName string, view View) *mux.Route {
	if view == nil {
		view = viewSimpleTemplate
	}
	handlerCtx := s.makeHandlerCtx(templateName)
	handler := view(handlerCtx)
	return s.mux.HandleFunc(pattern, handler)
}

func viewSimpleTemplate(ctx *HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.Template.Render(w, nil)
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
