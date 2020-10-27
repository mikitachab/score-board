package server

func (s *Server) setupRoutes() {
	s.Handle("/players", "players_list.html", handlePlayersList).Methods("GET")
	s.Handle("/players/add", "player_add.html", handleAddPlayer).Methods("GET", "POST")
	s.Handle("/", "index.html", nil).Methods("GET")
}
