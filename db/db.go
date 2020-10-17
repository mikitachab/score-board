package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(
		&Player{},
		&Play{},
		&SevenWondersScore{},
		&PlayScore{},
	)

	return db
}

func DBtest() {
	// TODO{mikitachab} this function should be removed in next PR
	db := GetDB()

	player1 := Player{Name: "Mikita"}
	player2 := Player{Name: "Irina"}
	play := Play{Date: time.Now(), Place: "Cyber"}
	score := SevenWondersScore{WonderName: "rhodos",
		MilitaryConflicts:    10,
		Wonders:              10,
		TreasuryContents:     10,
		CivilianStructures:   10,
		CommercialStructures: 10,
		Guilds:               10,
		ScientificStructures: 10}
	score2 := SevenWondersScore{WonderName: "babylon",
		MilitaryConflicts:    1,
		Wonders:              1,
		TreasuryContents:     1,
		CivilianStructures:   1,
		CommercialStructures: 1,
		Guilds:               1,
		ScientificStructures: 1}

	db.Create(&player1)
	db.Create(&player2)

	db.Create(&play)

	db.Create(&score)
	db.Create(&score2)

	db.Create(&PlayScore{Play: play,
		Player: player1, SevenWondersScore: score})
	db.Create(&PlayScore{Play: play,
		Player: player2, SevenWondersScore: score2})

	var playScores []PlayScore

	db.Joins("Player").Joins("Play").Joins("SevenWondersScore").Find(&playScores)

	for _, ps := range playScores {
		fmt.Println(ps.Player.Name, ps.SevenWondersScore.GetScoreSum())
	}
}
