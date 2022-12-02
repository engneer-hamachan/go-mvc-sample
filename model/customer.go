package model

import (
	"gorm.io/gorm"
	sqlite "main/config/database"
)

type Customer struct {
	gorm.Model
	Name string
	Age  int
}

func (c *Customer) Create() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	db.Create(c)

	connect.Close()
}

func (c *Customer) Update() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	db.Save(c)

	connect.Close()
}

func (c *Customer) Delete() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	db.Delete(c)

	connect.Close()
}
