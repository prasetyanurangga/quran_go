package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open("sqlite3", "quran.db")

	if err != nil {
		panic("Failed to Connect to Database")
	}

	DB = database
}