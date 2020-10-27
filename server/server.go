package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/mikitachab/score-board/db"
	"github.com/mikitachab/score-board/templateloader"
)

type Server struct {
	mux *mux.Router
	tl  *templateloader.TemplateLoader
	db  *gorm.DB
}

func NewServer() *Server {
	s := &Server{
		mux: mux.NewRouter(),
		tl:  templateloader.NewTemplateLoader(),
		db:  db.GetDB(),
	}

	s.setupRoutes()
	return s
}

func (s *Server) ListenAndServe(port string) error {
	return http.ListenAndServe(port, s.mux)
}

func handleErr(err error, message ...string) {
	if err != nil {
		if len(message) > 0 {
			err = fmt.Errorf("[%s] -- %w --", message[0], err)
		}
		log.Fatal(err)
	}
}

type TemplateHandlerFunc func(HandlerCtx) http.HandlerFunc

type HandlerCtx struct {
	S                  *Server
	RenderTemplateFunc templateloader.RenderTemplateFunc
}

func (s *Server) templateHandlerWrap(templateName string, handler TemplateHandlerFunc) http.HandlerFunc {
	rtf, err := s.tl.GetRenderTemplateFunc(templateName)
	handleErr(err, fmt.Sprintf("failed to setup %s template", templateName))
	return handler(HandlerCtx{s, rtf})
}

func handleSimpleTemplate(ctx HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := ctx.RenderTemplateFunc(w, nil)
		handleErr(err)
	}
}

func (s *Server) Handle(pattern, templateName string, handler TemplateHandlerFunc) *mux.Route {
	if handler == nil {
		handler = handleSimpleTemplate
	}
	wrappedHandler := s.templateHandlerWrap(templateName, handler)
	return s.mux.HandleFunc(pattern, wrappedHandler)
}
