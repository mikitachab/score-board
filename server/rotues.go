package server

func (s *Server) setupRoutes() {
	s.Handle("/players", "players_list.html", handlePlayersList).Methods("GET")
	s.Handle("/", "index.html", handleIndex).Methods("GET")
}
