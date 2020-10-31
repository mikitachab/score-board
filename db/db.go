package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GetDB is a function that setup database connection
// and return Repository instance to handle db operation
func GetDB() *Repository {
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

	return &Repository{db}
}
