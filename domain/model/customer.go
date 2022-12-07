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

