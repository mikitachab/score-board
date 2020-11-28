package server

func (s *Server) setupRoutes() {
	s.Handle("/players", "players_list.html", handlePlayersList).Methods("GET")
	s.Handle("/players/add", "add_player.html", handleAddPlayer).Methods("GET", "POST")
	s.Handle("/score/add/select", "add_score_select_players.html", handleAddScoreSelectPlayers).Methods("GET", "POST")
	s.Handle("/score/add/", "add_score.html", handleAddScore).Methods("GET", "POST")
	s.Handle("/", "index.html", handleIndex).Methods("GET")
}
