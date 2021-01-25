package database

import (
	"gorm.io/gorm"
	"ismaeldf/golang-gobarber/models"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(
		&models.Appointment{},
		&models.User{},
	)
}
