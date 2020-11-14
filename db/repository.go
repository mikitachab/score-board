package db

import (
	"gorm.io/gorm"
)

// RepositoryInterface define interface for db repo
type RepositoryInterface interface {
	GetAllPlayers() []Player
	GetAllPlays() []Play
	GetPlayScoresForPlay(Play) []PlayScore
}

// Repository is a wrapper which encapsulate db access operation
type Repository struct {
	db *gorm.DB
}

// GetAllPlayers return all players from database
func (r *Repository) GetAllPlayers() []Player {
	var players []Player
	r.db.Find(&players)
	return players
}

// GetPlayScoresForPlay returns all play scores for play
func (r *Repository) GetPlayScoresForPlay(play Play) []PlayScore {
	var playScores []PlayScore
	r.db.Joins("Player").Joins("Play").Joins("SevenWondersScore").Where("play_id", play.ID).Find(&playScores)
	return playScores
}

// GetAllPlays returns all plays from database
func (r *Repository) GetAllPlays() []Play {
	var plays []Play
	r.db.Find(&plays)
	return plays
}
