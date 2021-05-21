package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("games.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db

}
