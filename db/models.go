package db

import (
	"time"

	"gorm.io/gorm"
)

// Player is a model that represent each
// unique player
type Player struct {
	gorm.Model
	Name string `gorm:"unique;not null"`

	PlayScores []PlayScore
}

// Play is a model that represent fact that
// game was played
type Play struct {
	gorm.Model
	Place string
	Date  time.Time

	PlayScores []PlayScore
}

// SevenWondersScore is a simple representation
// of 7 Wonders game score
// TODO{mikitachab} add wonders type enum
type SevenWondersScore struct {
	gorm.Model
	WonderName           string
	MilitaryConflicts    int
	Wonders              int
	TreasuryContents     int
	CivilianStructures   int
	ScientificStructures int
	CommercialStructures int
	Guilds               int

	PlayScores []PlayScore
}

// GetScoreSum compute and return sum of every score field
func (sws *SevenWondersScore) GetScoreSum() int {
	return sws.CivilianStructures +
		sws.ScientificStructures +
		sws.MilitaryConflicts +
		sws.Wonders +
		sws.TreasuryContents +
		sws.CommercialStructures +
		sws.Guilds
}

// PlayScore is a junction table that
// join together information about
// play, player and its score
type PlayScore struct {
	gorm.Model
	PlayerID            uint
	Player              Player
	PlayID              uint
	Play                Play
	SevenWondersScoreID uint
	SevenWondersScore   SevenWondersScore
}
