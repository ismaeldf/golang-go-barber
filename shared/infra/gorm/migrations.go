package gorm

import (
	"gorm.io/gorm"
	entities2 "ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(
		&entities2.Appointment{},
		&entities.User{},
	)
}
