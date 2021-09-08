package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase(dbTarget string) {
	database, err := gorm.Open("sqlite3", dbTarget)
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Campaign{})

	DB = database
}
