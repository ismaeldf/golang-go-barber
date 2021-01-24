package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func CreateConnectionDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=gobarber port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	RunMigrations(db)

	return db
}