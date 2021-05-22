package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB returns a gorm.DB instance to access the database
func ConnectDB() *gorm.DB {

	// Open connection to the database 'games.db'
	db, err := gorm.Open(sqlite.Open("games.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db

}
