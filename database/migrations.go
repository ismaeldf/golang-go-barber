package database

import (
	"gorm.io/gorm"
	"ismaeldf.melo/golang/go-barber/models"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(&models.Appointment{})
}
