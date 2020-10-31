package db

import (
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func (r *DBRepository) GetAllPlayers() []Player {
	var players []Player
	r.db.Find(&players)
	return players
}
