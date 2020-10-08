package server

import "net/http"

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

func (s *Server) handleIndex() http.HandlerFunc {

	// you can put handler setup code here

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("<h1>Hello server</h1>"))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
