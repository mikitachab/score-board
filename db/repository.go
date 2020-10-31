package db

import (
	"gorm.io/gorm"
)

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
