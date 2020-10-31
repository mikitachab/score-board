package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() *DBRepository {
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

	return &DBRepository{db}
}
