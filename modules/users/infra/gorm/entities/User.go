package entities

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        string    `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"notnull" validate:"required"`
	Email     string    `json:"email" gorm:"notnull;unique" validate:"required"`
	Password  string    `json:"-" gorm:"notnull" validate:"required"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *User) isValid() error {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		fmt.Errorf("Error during User validation: %s", err.Error())
		return err
	}
	return nil
}


type UserUnhide struct {
	Password  string    `json:"password" gorm:"notnull" validate:"required"`
	User
}

func NewUser(name string, email string, password string) (*User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		Id:       uuid.NewV4().String(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		CreatedAt: time.Now(),
	}

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
