package db

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name string `gorm:"unique;not null"`

	PlayScores []PlayScore
}

type Play struct {
	gorm.Model
	Place string
	Date  time.Time

	PlayScores []PlayScore
}

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

func (sws *SevenWondersScore) GetScoreSum() int {
	return sws.CivilianStructures +
		sws.ScientificStructures +
		sws.MilitaryConflicts +
		sws.Wonders +
		sws.TreasuryContents +
		sws.CommercialStructures +
		sws.Guilds
}

type PlayScore struct {
	gorm.Model
	PlayerID            uint
	Player              Player
	PlayID              uint
	Play                Play
	SevenWondersScoreID uint
	SevenWondersScore   SevenWondersScore
}
