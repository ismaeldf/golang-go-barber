package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"notnull"`
	Email     string    `json:"email" gorm:"notnull;unique"`
	Password  string    `json:"password" gorm:"notnull"`
	CreatedAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
