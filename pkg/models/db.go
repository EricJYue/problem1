package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitDB() {
	var err error
	Db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	if err := Db.AutoMigrate(User{}); err != nil {
		log.Panic(err)
	}
}
