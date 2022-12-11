package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerId string
	Name       string
	Age        int
}

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mvc.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Customer{})

	return db
}
