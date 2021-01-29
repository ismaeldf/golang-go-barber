package gorm

import (
	"gorm.io/gorm"
	entitiesAppointments "ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	entitiesUsers "ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(
		&entitiesUsers.User{},
		&entitiesAppointments.Appointment{},
	)
}
