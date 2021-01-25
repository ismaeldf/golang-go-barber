package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"notnull"`
	Email     string    `json:"email" gorm:"notnull;unique"`
	Password  string    `json:"-" gorm:"notnull"`
	CreatedAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewUser(name string, email string, password string) *User{
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		ID: uuid.NewV4().String(),
		Name: name,
		Email: email,
		Password: string(hashedPassword),
	}
}