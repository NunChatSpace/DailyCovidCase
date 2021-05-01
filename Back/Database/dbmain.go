package Database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetupDatabase() {
	dbTemp, err := gorm.Open(sqlite.Open("mainDatabase.db"), &gorm.Config{})
	if err != nil {
		panic("Fail to connect to database")
	}

	db = dbTemp
}
