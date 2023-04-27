package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type UserModel struct {
	Id        uuid.UUID `json:"id" gorm:"not null;type:uuid"`
	Username  string    `json:"username" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
